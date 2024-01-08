package services

import (
	"Tuneless-Treasures/models"

	"gorm.io/gorm"
)

func ListOrders(db *gorm.DB) ([]models.Order, error) {
	orders := []models.Order{}
	err := db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrder(db *gorm.DB, id string) (*models.Order, error) {
	order := models.Order{}
	err := db.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func CreateOrder(db *gorm.DB, payload *models.OrderInput) (*models.Order, error) {
	var items []string
	if payload.Items == nil {
		items = []string{}
	} else {
		items = payload.Items
	}
	var total float64
	if payload.Total != 0 {
		total = payload.Total
	} else {
		var missingInstruments error 
		total, missingInstruments = calculateTotal(db, items)
		if missingInstruments != nil {
			return nil, missingInstruments
		}
		err := generateWorkOrderfForNewOrder(db, items)
		if err != nil {
			return nil, err
		}

	}
	order := models.Order{CustomerId: payload.CustomerId, Items: items, Total: total, Status: "created"}
	err := db.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func generateWorkOrderfForNewOrder(db *gorm.DB, items []string) error {
	itemsWithNoWorkOrders := []string{}
	for _, item := range items {
		instrument, err := getInstrument(db, item)
		if err != nil {
			itemsWithNoWorkOrders = append(itemsWithNoWorkOrders, item)
			return err
		}
		_, err = createWorkOrder(db, &models.WorkOrderInput{InstrumentID: instrument.ID})
		if err != nil {
			itemsWithNoWorkOrders = append(itemsWithNoWorkOrders, item)
			return err
		}
	}
	if len(itemsWithNoWorkOrders) > 0 {
		return models.ErrOrderWasCreatedWithMissingWorkOrders
	}
	return nil
}

func calculateTotal(db *gorm.DB, items []string) (float64, error) {
	total := 0.0
	missingInstruments := []string{}
	for _, item := range items {
		instrument, err := getInstrument(db, item)
		if err != nil {
			missingInstruments = append(missingInstruments, item)
		} else {
			total += instrument.Price
		}
	}
	if len(missingInstruments) > 0 {
		return 0, models.ErrInstrumentNotFound
	}
	return total, nil
}

func updateOrderStatus(db *gorm.DB, id string, status string) (*models.Order, error) {
	order, err := GetOrder(db, id)
	if err != nil {
		return nil, err
	}

	order.Status = status
	err = db.Save(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func completeOrder(db *gorm.DB, id string) (*models.Order, error) {
	return updateOrderStatus(db, id, "completed")
}

func orderProcessing(db *gorm.DB, id string) (*models.Order, error) {
	return updateOrderStatus(db, id, "processing")
}


func DeleteOrder(db *gorm.DB, id string) error {
	order, err := GetOrder(db, id)
	if err != nil {
		return err
	}

	err = db.Delete(&order).Error
	if err != nil {
		return err
	}

	return nil
}

func getCustomerOrders(db *gorm.DB, customerId string) ([]models.Order, error) {
	orders := []models.Order{}
	err := db.Where("customer_id = ?", customerId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func createInstrumentWorkOrders(db *gorm.DB, instruments []models.Instrument) error {
	for _, instrument := range instruments {
		_, err := createWorkOrder(db, &models.WorkOrderInput{InstrumentID: instrument.ID})
		if err != nil {
			return err
		}
		assignWorkerToWorkOrder(db, &models.WorkOrder{InstrumentID: instrument.ID}, 1)
	}
	return nil
}
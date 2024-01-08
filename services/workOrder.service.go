package services

import (
	"Tuneless-Treasures/models"

	"gorm.io/gorm"
)

func createWorkOrder(db *gorm.DB, payload *models.WorkOrderInput) (*models.WorkOrder, error) {
	workOrder := models.WorkOrder{InstrumentID: payload.InstrumentID, Status: "awaitingAssignment"}
	err := db.Create(&workOrder).Error
	if err != nil {
		return nil, err
	}
	assignemntErr := assignWorkOrder(db, &workOrder)
	if assignemntErr != nil {
		return nil, assignemntErr
	}

	return &workOrder, nil
}

func assignWorkOrder(db *gorm.DB, workOrder *models.WorkOrder) error {
	workOrder.Status = "assigned"
	instrumentCatrgory := workOrder.Instrument.CategoryID
	workerWithCategoryAsExpertise, err := getWorkerWithCategoryAsExpertise(db, instrumentCatrgory)
	if err != nil {
		return err
	}
	_, err = assignWorkerToWorkOrder(db, workOrder, workerWithCategoryAsExpertise.ID)
	if err != nil {
		return err
	}

	return nil
}

func getWorkerWithCategoryAsExpertise(db *gorm.DB, expertise uint) (*models.Worker, error) {
	worker := models.Worker{}
	err := db.Preload("Expertise").Where("expertise.category_id = ?", expertise).First(&worker).Error
	if err != nil {
		return nil, err
	}
	return &worker, nil
}

func assignWorkerToWorkOrder(db *gorm.DB, workOrder *models.WorkOrder, workerId uint) (*models.WorkOrder, error) {
	workOrder.WorkerId = workerId
	workOrder.Status = "assigned"
	err := db.Save(&workOrder).Error
	if err != nil {
		return nil, err
	}

	return workOrder, nil
}

func updateWorkOrderStatus(db *gorm.DB, workOrder *models.WorkOrder, status string) (*models.WorkOrder, error) {
	workOrder.Status = status
	err := db.Save(&workOrder).Error
	if err != nil {
		return nil, err
	}

	return workOrder, nil
}

func unassignWorkerFromWorkOrder(db *gorm.DB, workOrder *models.WorkOrder) (*models.WorkOrder, error) {
	workOrder.WorkerId = 0
	workOrder.Status = "awaitingAssignment"
	err := db.Save(&workOrder).Error
	if err != nil {
		return nil, err
	}

	return workOrder, nil
}
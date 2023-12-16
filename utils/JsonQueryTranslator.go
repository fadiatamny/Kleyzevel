package utils

import (
	"strings"

	"gorm.io/gorm"
)

func BuildCondition(gormQuery *gorm.DB, key string, value interface{}) *gorm.DB {
	// Handle MongoDB operators
	switch key {
	case "$and":
		// Handle $and operator
		andConditions, ok := value.([]interface{})
		if !ok {
			return gormQuery
		}
		for _, condition := range andConditions {
			gormQuery = BuildCondition(gormQuery, "", condition)
		}
	case "$or":
		// Handle $or operator
		orConditions, ok := value.([]interface{})
		if !ok {
			return gormQuery
		}
		orQuery := gormQuery
		for _, condition := range orConditions {
			orQuery = BuildCondition(orQuery, "", condition).Or(gormQuery)
		}
		gormQuery = orQuery
	default:
		// Handle other conditions
		if strings.Contains(key, ".") {
			// Handle nested fields
			gormQuery = gormQuery.Where(key+" = ?", value)
		} else {
			// Handle regular fields
			if fieldOperator, ok := value.(map[string]interface{}); ok {
				// Check for specific operators
				for operator, operand := range fieldOperator {
					switch operator {
					case "$in":
						gormQuery = gormQuery.Where(key+" IN (?)", operand)
						break
					case "$nin":
						gormQuery = gormQuery.Where(key+" NOT IN (?)", operand)
						break
					case "$ne":
						gormQuery = gormQuery.Where(key+" <> ?", operand)
						break
					case "$eq":
						gormQuery = gormQuery.Where(key+" = ?", operand)
						break
					case "$gt":
						gormQuery = gormQuery.Where(key+" > ?", operand)
						break
					case "$gte":
						gormQuery = gormQuery.Where(key+" >= ?", operand)
						break
					case "$lt":
						gormQuery = gormQuery.Where(key+" < ?", operand)
						break
					case "$lte":
						gormQuery = gormQuery.Where(key+" <= ?", operand)
						break
					case "$exists":
						// Example: {"field1": {"$exists": true}}
						exists, existsOk := operand.(bool)
						if existsOk && exists {
							gormQuery = gormQuery.Where(key + " IS NOT NULL")
						} else {
							gormQuery = gormQuery.Where(key + " IS NULL")
						}
						break
					// Add more cases for other operators as needed
					default:
						// Handle other operators or regular values
						gormQuery = gormQuery.Where(key+" = ?", value)
						break
					}
				}
			} else {
				// Handle regular values
				gormQuery = gormQuery.Where(key+" = ?", value)
			}
		}
	}
	return gormQuery
}

package main

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/samber/lo"
)

func removeRelationFromRecord[V comparable](record *models.Record, fieldName string, value V) *models.Record {
	projRecData := record.Data()
	resources, _ := projRecData[fieldName].([]V)
	newResources := lo.Without(resources, value)

	record.SetDataValue(fieldName, newResources)

	return record
}

func extendRelationFromRecord[V comparable](record *models.Record, fieldName string, value V) *models.Record {
	projRecData := record.Data()
	resources, _ := projRecData[fieldName].([]V)
	newResources := lo.Union(resources, []V{value})

	record.SetDataValue(fieldName, newResources)

	return record
}

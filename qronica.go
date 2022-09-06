package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type QronicaInstance struct {
	app *pocketbase.PocketBase

	resourcesCollection *models.Collection
	ProjectsCollection  *models.Collection
	StudiesCollection   *models.Collection
}

func NewQronica(app *pocketbase.PocketBase) (*QronicaInstance, error) {
	return &QronicaInstance{
		app: app,
		// principalDAO:        dao,
		// ResourcesCollection: resources,
		// ProjectsCollection:  projects,
		// StudiesCollection:   studies,
	}, nil
}

func (qi *QronicaInstance) ResourcesCollection(dao *daos.Dao) *models.Collection {
	if qi.resourcesCollection == nil {
		resources, err := dao.FindCollectionByNameOrId("resources")
		if err != nil {
			log.Println("Resources collection not found", err)
		}
		qi.resourcesCollection = resources
	}

	return qi.resourcesCollection
}

// func (qi *QronicaInstance) HydrateCollections() {
// 	log.Println(dao)

// 	projects, err := dao.FindCollectionByNameOrId("projects")
// 	if err != nil {
// 		log.Println("Projects collection not found", err)
// 	}

// 	studies, err := dao.FindCollectionByNameOrId("studies")
// 	if err != nil {
// 		log.Println("Studies collection not found", err)
// 	}

// 	qi.ResourcesCollection = resources
// 	qi.ProjectsCollection = projects
// 	qi.StudiesCollection = studies
// }

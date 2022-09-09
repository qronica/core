package main

import (
	"log"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type QronicaInstance struct {
	app *pocketbase.PocketBase

	resourcesCollection *models.Collection
	projectsCollection  *models.Collection
	// studiesCollection   *models.Collection

	projectStamps  map[string]ProjectStamp // hash table to save the last state of updated projects
	resourceStamps map[string]ResourceStamp
}

type UpdateRecordEventKind string

const (
	BeforeEvent UpdateRecordEventKind = "before"
	AfterEvent  UpdateRecordEventKind = "after"
)

type ProjectStamp struct {
	at        time.Time
	resources []string
}

type ResourceStamp struct {
	at       time.Time
	projects []string
}

func NewQronica(app *pocketbase.PocketBase) (*QronicaInstance, error) {
	return &QronicaInstance{
		app:            app,
		projectStamps:  map[string]ProjectStamp{},
		resourceStamps: map[string]ResourceStamp{},
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

func (qi *QronicaInstance) ProjectsCollection(dao *daos.Dao) *models.Collection {
	if qi.projectsCollection == nil {
		projects, err := dao.FindCollectionByNameOrId("projects")
		if err != nil {
			log.Println("Projects collection not found", err)
		}
		qi.projectsCollection = projects
	}

	return qi.projectsCollection
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

package main

import (
	"log"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/samber/lo"
)

func (qi *QronicaInstance) SideEffectAtCreateProject(dao *daos.Dao, e *core.RecordCreateEvent) error {
	data := e.Record.Data()
	resources, _ := data["resources"].([]string)

	projectID := e.Record.Id

	// for each resource, add the project to the list of its owns projects
	for _, resID := range resources {
		if resID == "" {
			log.Println("Resource ID is empty")
			continue
		}

		resource, err := dao.FindRecordById(qi.ResourcesCollection(dao), resID, nil)
		if err != nil {
			log.Println("Resource not found")
			continue
		}

		resource = extendRelationFromRecord(resource, "projects", projectID)
		if err := dao.SaveRecord(resource); err != nil {
			log.Println("Resource update failed")
			continue
		}
	}

	return nil
}

func (qi *QronicaInstance) SideEffectAtUpdateProject(kind UpdateRecordEventKind, dao *daos.Dao, e *core.RecordUpdateEvent) error {
	projectID := e.Record.Id

	oldProject, err := dao.FindRecordById(qi.ProjectsCollection(dao), projectID, nil)
	if err != nil {
		log.Println("Resource not found")
		return nil
	}

	resources, _ := oldProject.Data()["resources"].([]string)

	if kind == BeforeEvent {
		// save the stamp
		qi.projectStamps[projectID] = ProjectStamp{
			at:        time.Now(),
			resources: resources,
		}
	} else if kind == AfterEvent {
		old, exists := qi.projectStamps[projectID]
		if !exists {
			log.Println("old record doesn't exists")
			return nil
		}

		removedResources, newResources := lo.Difference(old.resources, resources)

		for _, resID := range removedResources {
			resource, err := dao.FindRecordById(qi.ResourcesCollection(dao), resID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			resource = removeRelationFromRecord(resource, "projects", projectID)

			if err := dao.SaveRecord(resource); err != nil {
				log.Println("Resource update failed")
				continue
			}
		}

		for _, resID := range newResources {
			resource, err := dao.FindRecordById(qi.ResourcesCollection(dao), resID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			resource = extendRelationFromRecord(resource, "projects", projectID)

			if err := dao.SaveRecord(resource); err != nil {
				log.Println("Resource update failed")
				continue
			}
		}

		// clean the stamp
		delete(qi.projectStamps, projectID)
	}

	return nil
}

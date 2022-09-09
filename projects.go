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
	log.Println(data)

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

		resRecData := resource.Data()
		projects, _ := resRecData["projects"].([]string)
		log.Println(resRecData)

		projects = lo.Union(projects, []string{e.Record.Id})

		log.Printf("adding project with id '%s' to your resource '%s'", e.Record.Id, resID)
		resource.SetDataValue("projects", projects)

		if err := dao.Save(resource); err != nil {
			log.Println("Resource update failed")
			continue
		}
	}

	return nil
}

func (qi *QronicaInstance) SideEffectAtUpdateProject(kind UpdateRecordEventKind, dao *daos.Dao, e *core.RecordUpdateEvent) error {
	data := e.Record.Data()
	projectID := e.Record.Id
	resources, _ := data["resources"].([]string)

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

		_, removedResources := lo.Difference(old.resources, resources)

		for _, resID := range removedResources {
			resource, err := dao.FindRecordById(qi.ResourcesCollection(dao), resID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			resRecData := resource.Data()
			projects, _ := resRecData["projects"].([]string)

			newProjects := lo.Without(projects, projectID)

			resource.SetDataValue("projects", newProjects)

			if err := dao.Save(resource); err != nil {
				log.Println("Resource update failed")
				continue
			}
		}

		// clean the stamp
		delete(qi.projectStamps, projectID)
	}

	return nil
}

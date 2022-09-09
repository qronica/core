package main

import (
	"log"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/samber/lo"
)

func (qi *QronicaInstance) SideEffectAtNewResource(dao *daos.Dao, e *core.RecordCreateEvent) error {
	data := e.Record.Data()
	projects, _ := data["projects"].([]string)
	log.Println(data)

	// for each resource, add the project to the list of its owns projects
	for _, projectID := range projects {
		if projectID == "" {
			log.Println("Project ID is empty")
			continue
		}

		project, err := dao.FindRecordById(qi.ProjectsCollection(dao), projectID, nil)
		if err != nil {
			log.Println("Project not found")
			continue
		}

		projectRecData := project.Data()
		resources, _ := projectRecData["resources"].([]string)
		log.Println(projectRecData)

		resources = lo.Union(resources, []string{e.Record.Id})

		log.Printf("adding resource with id '%s' to your project '%s'", e.Record.Id, projectID)
		project.SetDataValue("resources", resources)

		if err := dao.Save(project); err != nil {
			log.Println("Project update failed")
			continue
		}
	}

	return nil
}

func (qi *QronicaInstance) SideEffectAtUpdateResource(kind UpdateRecordEventKind, dao *daos.Dao, e *core.RecordUpdateEvent) error {
	data := e.Record.Data()
	resourceID := e.Record.Id
	projects, _ := data["projects"].([]string)

	if kind == BeforeEvent {
		// save the stamp
		qi.resourceStamps[e.Record.Id] = ResourceStamp{
			at:       time.Now(),
			projects: projects,
		}
	} else if kind == AfterEvent {
		old, exists := qi.resourceStamps[resourceID]
		if !exists {
			log.Println("old record doesn't exists")
			return nil
		}

		_, removedProjects := lo.Difference(old.projects, projects)

		for _, projID := range removedProjects {
			project, err := dao.FindRecordById(qi.ProjectsCollection(dao), projID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			projRecData := project.Data()
			resources, _ := projRecData["resources"].([]string)

			newResources := lo.Without(resources, resourceID)

			project.SetDataValue("resources", newResources)

			if err := dao.Save(project); err != nil {
				log.Println("Resource update failed")
				continue
			}
		}

		// clean the stamp
		delete(qi.resourceStamps, resourceID)
	}

	return nil
}

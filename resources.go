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
	resourceID := e.Record.Id

	oldResource, err := dao.FindRecordById(qi.ResourcesCollection(dao), resourceID, nil)
	if err != nil {
		log.Println("Resource not found")
		return nil
	}

	projects, _ := oldResource.Data()["projects"].([]string)

	if kind == BeforeEvent {
		// save the stamp
		log.Printf("saved stamp for resource '%s' with projects %v", resourceID, projects)
		qi.resourceStamps[resourceID] = ResourceStamp{
			at:       time.Now(),
			projects: projects,
		}
	} else if kind == AfterEvent {
		old, exists := qi.resourceStamps[resourceID]
		if !exists {
			log.Println("old record doesn't exists")
			return nil
		}

		log.Printf("recover stamp for resource '%s' with projects %v and old projects %v ", resourceID, projects, old.projects)

		// lo.Without(projects)
		removedProjects, newProjects := lo.Difference(old.projects, projects)

		log.Println(newProjects)
		log.Println(removedProjects)

		for _, projID := range removedProjects {
			project, err := dao.FindRecordById(qi.ProjectsCollection(dao), projID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			// projRecData := project.Data()
			// resources, _ := projRecData["resources"].([]string)

			// newResources := lo.Without(resources, resourceID)

			// project.SetDataValue("resources", newResources)
			project = removeRelationFromRecord(project, "resources", resourceID)

			if err := dao.SaveRecord(project); err != nil {
				log.Println("Resource update failed")
				continue
			}
		}

		for _, projID := range newProjects {
			project, err := dao.FindRecordById(qi.ProjectsCollection(dao), projID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			// projRecData := project.Data()
			// resources, _ := projRecData["resources"].([]string)

			// newResources := lo.Union(resources, []string{resourceID})

			// project.SetDataValue("resources", newResources)

			project = extendRelationFromRecord(project, "resources", resourceID)

			if err := dao.SaveRecord(project); err != nil {
				log.Println("Resource update failed")
				continue
			}
		}

		// clean the stamp
		delete(qi.resourceStamps, resourceID)
	}

	return nil
}

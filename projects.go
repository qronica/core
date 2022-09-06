package main

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/samber/lo"
)

func (qi *QronicaInstance) extendProjectsOfResourceWithNewProject(initialProjects []string, newProject string) []string {
	if lo.IndexOf(initialProjects, newProject) == -1 {
		return append(initialProjects, newProject)
	}

	return initialProjects
}

func (qi *QronicaInstance) SideEffectAtNewProject(dao *daos.Dao, e *core.RecordCreateEvent) error {
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

		projects = qi.extendProjectsOfResourceWithNewProject(projects, e.Record.Id)

		log.Printf("adding project with id '%s'", e.Record.Id)
		resource.SetDataValue("projects", projects)

		if err := dao.Save(resource); err != nil {
			log.Println("Resource update failed")
			continue
		}
	}

	return nil
}

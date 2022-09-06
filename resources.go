package main

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
)

// func (qi *QronicaInstance) extendProjectsOfResourceWithNewProject(initialProjects []string, newProject string) []string {
// 	if lo.IndexOf(initialProjects, newProject) == -1 {
// 		return append(initialProjects, newProject)
// 	}

// 	return initialProjects
// }

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

		resources = qi.extendProjectsOfResourceWithNewProject(resources, e.Record.Id)

		log.Printf("adding resource with id '%s' to your project '%s'", e.Record.Id, projectID)
		project.SetDataValue("resources", resources)

		if err := dao.Save(project); err != nil {
			log.Println("Project update failed")
			continue
		}
	}

	return nil
}

func (qi *QronicaInstance) SideEffectAtUpdateResource(dao *daos.Dao, e *core.RecordUpdateEvent) error {
	return qi.SideEffectAtNewResource(dao, &core.RecordCreateEvent{Record: e.Record})
}

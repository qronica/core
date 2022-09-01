package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	_ "github.com/qronica/core/migrations"
)

func main() {
	app := pocketbase.New()

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		log.Printf("Collection name: %s", e.Record.Collection().Name) // still unsaved

		if e.Record.Collection().Name != "projects" {
			log.Println("Not a project")
			return nil
		}

		data := e.Record.Data()
		resources, _ := data["resources"].([]string)
		log.Println(data)

		for _, resID := range resources {
			if resID == "" {
				log.Println("Resource ID is empty")
				continue
			}

			resColl, err := app.Dao().FindCollectionByNameOrId("resources")
			if err != nil {
				log.Println("Resource collection not found")
				continue
			}
			// var resource :=
			resRec, err := app.Dao().FindRecordById(resColl, resID, nil)
			if err != nil {
				log.Println("Resource not found")
				continue
			}

			resRecData := resRec.Data()
			project, _ := resRecData["project"].(string)

			log.Println(resRecData)

			if project == "" {
				log.Println("Project not found")
				// continue
			}

			log.Printf("setting project with id '%s'", e.Record.Id)
			resRec.SetDataValue("project", e.Record.Id)

			if err := app.Dao().Save(resRec); err != nil {
				log.Println("Resource update failed")
				continue
			}

		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

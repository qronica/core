package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	_ "github.com/qronica/core/migrations"
)

func main() {
	app := pocketbase.New()

	qronica, err := NewQronica(app)
	if err != nil {
		log.Fatal(err)
	}

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		collectionName := e.Record.Collection().Name
		log.Printf("Collection name: %s", collectionName) // still unsaved

		if collectionName == "projects" {
			return qronica.SideEffectAtNewProject(app.Dao(), e)
		} else if collectionName == "resources" {
			return qronica.SideEffectAtNewResource(app.Dao(), e)
		}

		return nil
	})

	app.OnRecordAfterUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		collectionName := e.Record.Collection().Name
		log.Printf("Collection name: %s", collectionName) // still unsaved

		if collectionName == "projects" {
			return qronica.SideEffectAtUpdateProject(app.Dao(), e)
		} else if collectionName == "resources" {
			return qronica.SideEffectAtUpdateResource(app.Dao(), e)
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

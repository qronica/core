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
			return qronica.SideEffectAtNewProject(e, app.Dao())
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

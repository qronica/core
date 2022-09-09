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

	// Projects

	app.OnRecordAfterCreateRequest().Add(func(data *core.RecordCreateEvent) error {
		log.Println("OnRecordAfterCreateRequest for [%s] %s", data.Record.Collection().Name, data.Record.Id)

		if data.Record.Collection().Name == "projects" {
			return qronica.SideEffectAtCreateProject(app.Dao(), data)
		}

		return nil
	})

	app.OnRecordBeforeUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "projects" {
			return qronica.SideEffectAtUpdateProject(BeforeEvent, app.Dao(), data)
		}

		return nil
	})

	app.OnRecordAfterUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "projects" {
			return qronica.SideEffectAtUpdateProject(AfterEvent, app.Dao(), data)
		}

		return nil
	})

	// Resources

	app.OnRecordAfterCreateRequest().Add(func(data *core.RecordCreateEvent) error {
		if data.Record.Collection().Name == "resources" {
			return qronica.SideEffectAtNewResource(app.Dao(), data)
		}

		return nil
	})

	app.OnRecordBeforeUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "resources" {
			return qronica.SideEffectAtUpdateResource(BeforeEvent, app.Dao(), data)
		}

		return nil
	})

	app.OnRecordAfterUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "resources" {
			return qronica.SideEffectAtUpdateResource(AfterEvent, app.Dao(), data)
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

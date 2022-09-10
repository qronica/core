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

	app.OnRecordAfterCreateRequest().Add(func(data *core.RecordCreateEvent) error {
		if data.Record.Collection().Name == "projects" {
			// log.Printf("SideEffectAtCreateProject for [%s] %s", data.Record.Collection().Name, data.Record.Id)
			return qronica.SideEffectAtCreateProject(app.Dao(), data)
		}

		return nil
	})

	app.OnRecordAfterCreateRequest().Add(func(data *core.RecordCreateEvent) error {
		if data.Record.Collection().Name == "resources" {
			// log.Printf("SideEffectAtNewResource for [%s] %s", data.Record.Collection().Name, data.Record.Id)
			return qronica.SideEffectAtNewResource(app.Dao(), data)
		}

		return nil
	})

	app.OnRecordBeforeUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "projects" {
			// log.Printf("SideEffectAtUpdateProject(BeforeEvent) for [%s] %s", data.Record.Collection().Name, data.Record.Id)
			return qronica.SideEffectAtUpdateProject(BeforeEvent, app.Dao(), data)
		}

		return nil
	})

	app.OnRecordBeforeUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "resources" {
			// log.Printf("SideEffectAtUpdateResource(Before) for [%s] %s", data.Record.Collection().Name, data.Record.Id)
			// log.Printf("%+v", data.Record.Data())
			return qronica.SideEffectAtUpdateResource(BeforeEvent, app.Dao(), data)
		}

		return nil
	})

	app.OnRecordAfterUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "projects" {
			// log.Printf("SideEffectAtUpdateProject(After) for [%s] %s", data.Record.Collection().Name, data.Record.Id)
			return qronica.SideEffectAtUpdateProject(AfterEvent, app.Dao(), data)
		}

		return nil
	})

	app.OnRecordAfterUpdateRequest().Add(func(data *core.RecordUpdateEvent) error {
		if data.Record.Collection().Name == "resources" {
			// log.Printf("SideEffectAtUpdateResource(After) for [%s] %s", data.Record.Collection().Name, data.Record.Id)
			// log.Printf("%+v", data.Record.Data())
			return qronica.SideEffectAtUpdateResource(AfterEvent, app.Dao(), data)
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

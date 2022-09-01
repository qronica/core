package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "systemprofiles0",
				"created": "2022-09-01 18:57:22.069",
				"updated": "2022-09-01 18:57:22.069",
				"name": "profiles",
				"system": true,
				"schema": [
					{
						"system": true,
						"id": "pbfielduser",
						"name": "userId",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "pbfieldname",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pbfieldavatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "userId = @request.user.id",
				"viewRule": "userId = @request.user.id",
				"createRule": "userId = @request.user.id",
				"updateRule": "userId = @request.user.id",
				"deleteRule": null
			},
			{
				"id": "lvttbwx1jksy9nz",
				"created": "2022-09-01 19:01:25.230",
				"updated": "2022-09-01 19:05:07.235",
				"name": "projects",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "bazpzkd7",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "weds4nhk",
						"name": "scope",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"private",
								"internal",
								"public"
							]
						}
					},
					{
						"system": false,
						"id": "vhrpbief",
						"name": "owner",
						"type": "user",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "x8pmyhcs",
						"name": "space",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": "",
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}

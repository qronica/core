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
				"id": "lvttbwx1jksy9nz",
				"created": "2022-09-01 19:01:25.230",
				"updated": "2022-09-09 01:47:13.416",
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
					},
					{
						"system": false,
						"id": "570h38my",
						"name": "resources",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 9999,
							"collectionId": "rwskiv5gp4socah",
							"cascadeDelete": false
						}
					}
				],
				"listRule": "",
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "rwskiv5gp4socah",
				"created": "2022-09-01 19:16:38.039",
				"updated": "2022-09-09 01:49:50.608",
				"name": "resources",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "q0rrwjoo",
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
						"id": "9itgclqj",
						"name": "kind",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"document",
								"text",
								"sound",
								"image",
								"video"
							]
						}
					},
					{
						"system": false,
						"id": "qa5b93mr",
						"name": "file",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [],
							"thumbs": [
								"0x100"
							]
						}
					},
					{
						"system": false,
						"id": "fr14acfo",
						"name": "text",
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
						"id": "heklyfxk",
						"name": "link",
						"type": "url",
						"required": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "juemaxkp",
						"name": "projects",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99999,
							"collectionId": "lvttbwx1jksy9nz",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "hetxqzoy",
						"name": "parents",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99999,
							"collectionId": "rwskiv5gp4socah",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "ru18pot3",
						"name": "children",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99999,
							"collectionId": "rwskiv5gp4socah",
							"cascadeDelete": false
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "systemprofiles0",
				"created": "2022-09-09 01:47:13.414",
				"updated": "2022-09-09 01:47:13.416",
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

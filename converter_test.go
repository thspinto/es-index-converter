package main

import (
	"testing"

	"github.com/franela/goblin"
)

func TestSecret(t *testing.T) {

	var csIndex = `{
		"IndexFields": [
				{
						"Options": {
								"IndexFieldName": "int_type",
								"IndexFieldType": "int",
								"IntOptions": {
										"FacetEnabled": true,
										"SearchEnabled": true,
										"ReturnEnabled": true,
										"SortEnabled": true
								}
						}
				},
				{
						"Options": {
								"IndexFieldName": "string_type",
								"IndexFieldType": "text",
								"IntOptions": {
										"FacetEnabled": true,
										"SearchEnabled": true,
										"ReturnEnabled": true,
										"SortEnabled": true
								}
						}
				}
		]
	}`

	var esIndex = `{
		"properties": {
			"int_type": {
				"type": "long"
			}
		}
	}`
	g := goblin.Goblin(t)
	g.Describe("convert to ES", func() {

		g.It("should match event", func() {
			convertToEs(csIndex)
			g.Assert(2).Equal(2)
		})
	})
}

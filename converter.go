package main

import (
	"encoding/json"
	"fmt"
)

type CsIndex struct {
	IndexFields []struct {
		Options struct {
			IndexFieldName string `json:"IndexFieldName"`
			IndexFieldType string `json:"IndexFieldType"`
			IntOptions     struct {
				FacetEnabled  bool `json:"FacetEnabled"`
				SearchEnabled bool `json:"SearchEnabled"`
				ReturnEnabled bool `json:"ReturnEnabled"`
				SortEnabled   bool `json:"SortEnabled"`
			} `json:"IntOptions"`
		} `json:"Options"`
	} `json:"IndexFields"`
}

type Mapping map[string]struct {
	Type string `json:"type"`
}

type EsIndex struct {
	properties []Mapping
}

func convertToEs(csIndex string) string {
	var index CsIndex
	json.Unmarshal([]byte(csIndex), &index)
	fmt.Println(index.IndexFields[0].Options.IndexFieldName)

	return ""
}

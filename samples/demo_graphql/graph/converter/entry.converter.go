package converter

import (
	"demo_graphql/graph/model"
	"encoding/json"
	"strings"
)

type (
	EntryConverter interface {
		ConvertToEntries(jsonInput string) []*model.Entry
		ConvertToJSON(entry *model.Entry) string
	}
	entryConverterImpl struct{}
)

func NewEntryConverter() EntryConverter {
	return &entryConverterImpl{}
}

func (c *entryConverterImpl) ConvertToEntries(jsonInput string) []*model.Entry {
	var entries []map[string]interface{}
	err := json.Unmarshal([]byte(jsonInput), &entries)
	if err != nil {
		return nil
	}

	var result []*model.Entry
	for _, e := range entries {
		entry := &model.Entry{
			ID:                       e["id"].(string),
			Title:                    e["title"].(string),
			Assignee:                 e["assignee"].(string),
			Author:                   strings.Split(e["inventor/author"].(string), ", "),
			PriorityDate:             e["priority date"].(string),
			CreationDate:             e["filing/creation date"].(string),
			PublicationDate:          e["publication date"].(string),
			GrantDate:                e["grant date"].(string),
			ResultLink:               e["result link"].(string),
			RepresentativeFigureLink: e["representative figure link"].(string),
		}
		result = append(result, entry)
	}

	return result
}

func (c *entryConverterImpl) ConvertToJSON(entry *model.Entry) string {
	jsonData, err := json.Marshal(entry)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

package services_test

import (
	"testing"

	"demo_graphql/graph/model"
	"demo_graphql/graph/services"
)

type (
	mockConverter struct{}
	mockRepository struct{}
)

var mockData []*model.Entry = []*model.Entry{
	{
		ID: "JP-5948372-B2",
		Title: "音声アシスタントでの以前のユーザインタラクションに基づく意図導出",
		Assignee: "アップル インコーポレイテッド, アップル  インコーポレイテッド",
		Author: "アダム ジョン チャイヤー, アンドリュー ジョン ハーキンズ",
		PriorityDate: "2010-01-18",
		CreationDate: "2014-06-20",
		PublicationDate: "2016-07-06",
		GrantDate: "2016-07-06",
		ResultLink: "https://patents.google.com/patent/JP5948372B2/ja",
		RepresentativeFigureLink: "",
	},
}

func (mc *mockConverter) ConvertToEntries(fileContent string) []*model.Entry {
	return mockData
}

func (mc *mockConverter) ConvertToJSON(entry *model.Entry) string {
	return ""
}

func (mr *mockRepository) Fetch() (string, error) {
	return "", nil
}

func TestGetEntryByID(t *testing.T) {
	service := services.NewEntryService(
		&mockConverter{},
		&mockRepository{},
	)

	tests := []struct {
		name    string
		id      string
		want    *model.Entry
	}{
		{"EntryExists", "JP-5948372-B2", mockData[0]},
		{"EntryDoesNotExist", "xxx", &model.Entry{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetEntryByID(tt.id)
			// 異常系はpass
			if err != nil {
				return
			}
			// 正しいEntryを返す
			if got != tt.want {
				t.Errorf("GetEntryByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

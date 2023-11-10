package services

import (
	"demo_graphql/graph/model"
)

type (
	EntryService interface {
		GetEntryByID(id string) (*model.Entry, error)
		SearchEntryByTitle(title string) ([]*model.Entry, error)
		AddEntry() (*model.Entry, error)
	}
	entryServiceImpl struct{}
)

func NewEntryService() EntryService {
	return &entryServiceImpl{}
}

func (s *entryServiceImpl) GetEntryByID(id string) (*model.Entry, error) {
	// test
	return &model.Entry{
		ID: "JP-5948372-B2",
		Title: "音声アシスタントでの以前のユーザインタラクションに基づく意図導出",
		Assignee: "アップル インコーポレイテッド, アップル  インコーポレイテッド",
		Author: []string{"アダム ジョン チャイヤー，, アダム  ジョン チャイヤー，, ディディエ ルネ グッツオーニ，, ディディエ  ルネ グッツオーニ，, トーマス ロバート グルーバー，, トーマス  ロバート グルーバー，, クリストファー ディーン ブリガム，, クリストファー  ディーン ブリガム，"},
		PriorityDate: "2010-01-18",
		CreationDate: "2014-06-20",
		PublicationDate: "2016-07-06",
		GrantDate: "2016-07-06",
		ResultLink: "https://patents.google.com/patent/JP5948372B2/ja",
		RepresentativeFigureLink: "",
	}, nil
}

func (s *entryServiceImpl) SearchEntryByTitle(title string) ([]*model.Entry, error) {
	return nil, nil
}

func (s *entryServiceImpl) AddEntry() (*model.Entry, error) {
	return nil, nil
}

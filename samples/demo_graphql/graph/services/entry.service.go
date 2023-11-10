package services

import (
	"demo_graphql/graph/converter"
	"demo_graphql/graph/model"
	"io/ioutil"
	"log"
)

type (
	EntryService interface {
		GetEntryByID(id string) (*model.Entry, error)
		SearchEntryByTitle(title string) ([]*model.Entry, error)
		AddEntry() (*model.Entry, error)
	}
	entryServiceImpl struct{
		converter converter.EntryConverter
	}
)

func NewEntryService(
	c converter.EntryConverter,
) EntryService {
	return &entryServiceImpl{
		converter: c,
	}
}

func (s *entryServiceImpl) GetEntryByID(id string) (*model.Entry, error) {
	fileContent, err := ioutil.ReadFile("data/entry.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	entries := s.converter.ConvertToEntries(string(fileContent))
	return entries[1], nil
}

func (s *entryServiceImpl) SearchEntryByTitle(title string) ([]*model.Entry, error) {
	return nil, nil
}

func (s *entryServiceImpl) AddEntry() (*model.Entry, error) {
	return nil, nil
}

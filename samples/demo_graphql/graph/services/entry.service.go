package services

import (
	"demo_graphql/graph/converter"
	"demo_graphql/graph/model"
	"errors"
	"fmt"
	"io/ioutil"
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
	// TODO: separate to repository
	fileContent, err := ioutil.ReadFile("data/entry.txt")
	if err != nil {
		panic(fmt.Errorf("Error reading file: %v", err))
	}

	entries := s.converter.ConvertToEntries(string(fileContent))

	for _, entry := range entries {
		if entry.ID == id {
			return entry, nil
		}
	}
	return &model.Entry{}, errors.New("entry not found")
}

func (s *entryServiceImpl) SearchEntryByTitle(title string) ([]*model.Entry, error) {
	return nil, nil
}

func (s *entryServiceImpl) AddEntry() (*model.Entry, error) {
	return nil, nil
}

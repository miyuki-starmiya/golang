package services

import (
	"errors"

	"demo_graphql/graph/converter"
	"demo_graphql/graph/model"
	"demo_graphql/graph/repositories"
)

type (
	EntryService interface {
		GetEntryByID(id string) (*model.Entry, error)
		SearchEntryByTitle(title string) ([]*model.Entry, error)
		AddEntry() (*model.Entry, error)
	}
	entryServiceImpl struct{
		converter converter.EntryConverter
		repository repositories.EntryRepository
	}
)

func NewEntryService(
	c converter.EntryConverter,
	r repositories.EntryRepository,
) EntryService {
	return &entryServiceImpl{
		converter: c,
		repository: r,
	}
}

func (s *entryServiceImpl) GetEntryByID(id string) (*model.Entry, error) {
	fileContent, err := s.repository.Fetch()
	if err != nil {
		return nil, err
	}

	entries := s.converter.ConvertToEntries(fileContent)

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

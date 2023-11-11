package repositories

import (
	"errors"
	"io/ioutil"
)

type (
	EntryRepository interface {
		Fetch() (string, error)
	}
	entryRepositoryImpl struct{}
)

func NewEntryRepository() EntryRepository {
	return &entryRepositoryImpl{}
}

func (r *entryRepositoryImpl) Fetch() (string, error) {
	fileContent, err := ioutil.ReadFile("data/entry.txt")
	if err != nil {
		return "", errors.New("Error reading file")
	}
	return string(fileContent), nil
}

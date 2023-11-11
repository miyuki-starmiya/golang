package services_test

import (
	"io/ioutil"
	"os"
	"testing"

	"demo_graphql/graph/model"
	"demo_graphql/graph/services"
)

type MockConverter struct{}

func (mc *MockConverter) ConvertToEntries(fileContent string) []*model.Entry {
	// Mock conversion logic
	return []*model.Entry{
		{ID: "1", /* other fields */},
		{ID: "2", /* other fields */},
	}
}

func (mc *MockConverter) ConvertToJSON(entry *model.Entry) string {
	return ""
}

func TestGetEntryByID(t *testing.T) {
	// Setup
	testData := `[{"id": "1", "name": "Test Entry 1"}, {"id": "2", "name": "Test Entry 2"}]`
	tempFile, err := ioutil.TempFile("", "entry")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Create an instance of entryServiceImpl with a fake converter
	service := services.NewEntryService(
		&MockConverter{},
	)

	// Test cases
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{"EntryExists", "1", false},
		{"EntryDoesNotExist", "3", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetEntryByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEntryByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

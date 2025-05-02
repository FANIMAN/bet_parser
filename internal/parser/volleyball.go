package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"betting-parser/internal/models"
)

// ParseVolleyballPrematch parses the volleyball prematch JSON data
func ParseVolleyballPrematch(filePath string) (models.VolleyballPrematchRoot, error) {
	var data models.VolleyballPrematchRoot

	// Read the JSON file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return data, fmt.Errorf("error reading file: %v", err)
	}

	// Parse the JSON data into Go struct
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return data, fmt.Errorf("error unmarshalling data: %v", err)
	}

	return data, nil
}

// ParseVolleyballResult parses the volleyball result JSON data
func ParseVolleyballResult(filePath string) (models.VolleyballResultRoot, error) {
	// Similar logic as for prematch parsing
	var data models.VolleyballResultRoot

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return data, fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return data, fmt.Errorf("error unmarshalling data: %v", err)
	}

	return data, nil
}

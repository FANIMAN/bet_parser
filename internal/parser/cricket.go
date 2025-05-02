package parser


// parser/cricket
import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"betting-parser/internal/models"
)

// ParseCricketPrematch parses the cricket prematch JSON data
func ParseCricketPrematch(filePath string) (models.CricketPrematchRoot, error) {
	var data models.CricketPrematchRoot

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

// ParseCricketResult parses the cricket result JSON data
func ParseCricketResult(filePath string) (models.CricketResultRoot, error) {
	// Similar logic as for prematch parsing
	var data models.CricketResultRoot

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

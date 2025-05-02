package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	_ = os.MkdirAll("testdata", os.ModePerm)

	files := map[string]interface{}{
		"testdata/cricket_prematch.json":     cricketPrematchData(),
		"testdata/cricket_result.json":       cricketResultData(),
		"testdata/volleyball_prematch.json":  volleyballPrematchData(),
		"testdata/volleyball_result.json":    volleyballResultData(),
	}

	for path, data := range files {
		if err := writeJSON(path, data); err != nil {
			fmt.Printf("Failed to write %s: %v\n", path, err)
		} else {
			fmt.Printf("✅ Wrote %s\n", path)
		}
	}
}

func writeJSON(path string, data interface{}) error {
	file, err := os.Create(filepath.Clean(path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func cricketPrematchData() map[string]interface{} {
	return map[string]interface{}{
		"success": 1,
		"result": []map[string]interface{}{
			{
				"id":          "10001",
				"time":        "2025-05-01 10:00:00",
				"time_status": "0",
				"league": map[string]interface{}{
					"id":   "1",
					"name": "World Cricket League",
					"cc":   "in",
				},
				"home": map[string]interface{}{
					"id":       "101",
					"name":     "India",
					"image_id": "1001",
					"cc":       "in",
				},
				"away": map[string]interface{}{
					"id":       "102",
					"name":     "Australia",
					"image_id": "1002",
					"cc":       "au",
				},
				"odds": map[string]interface{}{
					"pre_match": []map[string]interface{}{
						{
							"label": "Winner",
							"bookmakers": []map[string]interface{}{
								{
									"name": "Bet365",
									"odds": map[string]string{
										"1": "1.50",
										"2": "2.60",
									},
								},
							},
						},
					},
				},
				"headers": []map[string]interface{}{
					{
						"name": "Main",
						"markets": []map[string]interface{}{
							{
								"name": "Match Winner",
								"odds": []map[string]string{
									{"label": "India", "value": "1.50"},
									{"label": "Australia", "value": "2.60"},
								},
							},
							{
								"name": "Total Runs Over/Under",
								"odds": []map[string]string{
									{"label": "Over 200.5", "value": "1.90"},
									{"label": "Under 200.5", "value": "1.80"},
								},
							},
						},
					},
				},
			},
		},
	}
}

func cricketResultData() map[string]interface{} {
	return map[string]interface{}{
		"success": 1,
		"results": []map[string]interface{}{
			{
				"id":          "10001",
				"sport_id":    "3",
				"time":        "1746108000",
				"time_status": "3",
				"league": map[string]interface{}{
					"id":   "1",
					"name": "World Cricket League",
					"cc":   "in",
				},
				"home": map[string]interface{}{
					"id":       "101",
					"name":     "India",
					"image_id": "1001",
					"cc":       "in",
				},
				"away": map[string]interface{}{
					"id":       "102",
					"name":     "Australia",
					"image_id": "1002",
					"cc":       "au",
				},
				"ss": "220-210",
				"extra": map[string]interface{}{
					"stadium_data": map[string]interface{}{
						"id":           "10",
						"name":         "Eden Gardens",
						"city":         "Kolkata",
						"country":      "India",
						"capacity":     "66000",
						"googlecoords": "22.5645,88.3433",
					},
				},
				"has_lineup":        1,
				"inplay_created_at": "1746107172",
				"inplay_updated_at": "1746121199",
				"confirmed_at":      "1746122977",
				"bet365_id":         "987654",
			},
		},
	}
}

func volleyballPrematchData() map[string]interface{} {
	return map[string]interface{}{
		"success": 1,
		"result": []map[string]interface{}{
			{
				"id":          "20001",
				"time":        "2025-05-01 13:00:00",
				"time_status": "0",
				"league": map[string]interface{}{
					"id":   "2",
					"name": "World Volleyball Cup",
					"cc":   "us",
				},
				"home": map[string]interface{}{
					"id":       "201",
					"name":     "USA",
					"image_id": "2001",
					"cc":       "us",
				},
				"away": map[string]interface{}{
					"id":       "202",
					"name":     "Brazil",
					"image_id": "2002",
					"cc":       "br",
				},
				"odds": map[string]interface{}{
					"pre_match": []map[string]interface{}{
						{
							"label": "1X2",
							"bookmakers": []map[string]interface{}{
								{
									"name": "Bet365",
									"odds": map[string]string{
										"1": "1.80",
										"X": "3.20",
										"2": "2.10",
									},
								},
							},
						},
					},
				},
				"headers": []map[string]interface{}{
					{
						"name": "Main",
						"markets": []map[string]interface{}{
							{
								"name": "1X2",
								"odds": []map[string]string{
									{"label": "USA", "value": "1.80"},
									{"label": "Draw", "value": "3.20"},
									{"label": "Brazil", "value": "2.10"},
								},
							},
							{
								"name": "Correct Score",
								"odds": []map[string]string{
									{"label": "3-1", "value": "4.50"},
									{"label": "3-2", "value": "5.00"},
								},
							},
						},
					},
				},
			},
		},
	}
}

func volleyballResultData() map[string]interface{} {
	return map[string]interface{}{
		"success": 1,
		"results": []map[string]interface{}{
			{
				"id":          "20001",
				"sport_id":    "2",
				"time":        "1746115200",
				"time_status": "3",
				"league": map[string]interface{}{
					"id":   "2",
					"name": "World Volleyball Cup",
					"cc":   "us",
				},
				"home": map[string]interface{}{
					"id":       "201",
					"name":     "USA",
					"image_id": "2001",
					"cc":       "us",
				},
				"away": map[string]interface{}{
					"id":       "202",
					"name":     "Brazil",
					"image_id": "2002",
					"cc":       "br",
				},
				"ss": "3-2",
			},
		},
	}
}

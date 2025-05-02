package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Team struct {
	CC      string `json:"cc"`
	ID      string `json:"id"`
	ImageID string `json:"image_id"`
	Name    string `json:"name"`
}

type MarketOdd struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Market struct {
	Name string      `json:"name"`
	Odds []MarketOdd `json:"odds"`
}

type Header struct {
	Name    string   `json:"name"`
	Markets []Market `json:"markets"`
}

type Bookmaker struct {
	Name string            `json:"name"`
	Odds map[string]string `json:"odds"`
}

type OddsInfo struct {
	Label      string      `json:"label"`
	Bookmakers []Bookmaker `json:"bookmakers"`
}

type League struct {
	CC   string `json:"cc"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PreMatchGame struct {
	ID      string    `json:"id"`
	Home    Team      `json:"home"`
	Away    Team      `json:"away"`
	League  League    `json:"league"`
	Time    string    `json:"time"`
	Headers []Header  `json:"headers"`
	Odds    struct {
		PreMatch []OddsInfo `json:"pre_match"`
	} `json:"odds"`
	TimeStatus string `json:"time_status"`
}

type CricketPreMatch struct {
	Result  []PreMatchGame `json:"result"`
	Success int            `json:"success"`
}

type Stadium struct {
	Capacity     string `json:"capacity"`
	City         string `json:"city"`
	Country      string `json:"country"`
	GoogleCoords string `json:"googlecoords"`
	ID           string `json:"id"`
	Name         string `json:"name"`
}

type Extra struct {
	StadiumData Stadium `json:"stadium_data"`
}

type ResultGame struct {
	ID              string `json:"id"`
	Home            Team   `json:"home"`
	Away            Team   `json:"away"`
	League          League `json:"league"`
	TimeStatus      string `json:"time_status"`
	SS              string `json:"ss"` // Score string
	ConfirmedAt     string `json:"confirmed_at"`
	HasLineup       int    `json:"has_lineup"`
	InplayCreatedAt string `json:"inplay_created_at"`
	InplayUpdatedAt string `json:"inplay_updated_at"`
	Extra           Extra  `json:"extra"`
	SportID         string `json:"sport_id"`
	Time            string `json:"time"`
}

type CricketResult struct {
	Results []ResultGame `json:"results"`
	Success int          `json:"success"`
}

func readJSON[T any](filename string, data *T) error {
	absPath, _ := filepath.Abs(filename)
	content, err := os.ReadFile(absPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, data)
}

func main() {
	var prematch CricketPreMatch
	var result CricketResult

	// Read files
	if err := readJSON("testdata/cricket_prematch.json", &prematch); err != nil {
		fmt.Println("Failed to load pre-match:", err)
		return
	}
	if err := readJSON("testdata/cricket_result.json", &result); err != nil {
		fmt.Println("Failed to load result:", err)
		return
	}

	fmt.Println("Loaded pre-match and result data successfully.")

	if len(prematch.Result) == 0 || len(result.Results) == 0 {
		fmt.Println("No match data found.")
		return
	}

	game := prematch.Result[0]
	final := result.Results[0]

	fmt.Printf("Match: %s vs %s\n", game.Home.Name, game.Away.Name)
	fmt.Println("Evaluating selections...\n")

	scoreParts := strings.Split(final.SS, "-")
	if len(scoreParts) != 2 {
		fmt.Println("Invalid final score format:", final.SS)
		return
	}

	homeScore := scoreParts[0]
	awayScore := scoreParts[1]

	// Sample 1: Match Winner
	for _, header := range game.Headers {
		for _, market := range header.Markets {
			if market.Name == "Match Winner" {
				fmt.Println("Market: Match Winner")

				for _, odd := range market.Odds {
					fmt.Printf("- Selection: %s at odds %s\n", odd.Label, odd.Value)

					winTeam := ""
					if final.SS != "" {
						hs, as := toInt(homeScore), toInt(awayScore)
						if hs > as {
							winTeam = game.Home.Name
						} else if as > hs {
							winTeam = game.Away.Name
						}
					}

					if odd.Label == winTeam {
						fmt.Println("  ✅ Result: WON")
					} else {
						fmt.Println("  ❌ Result: LOST")
					}
				}
			}

			if strings.Contains(market.Name, "Over/Under") {
				fmt.Println("\nMarket:", market.Name)
				for _, odd := range market.Odds {
					fmt.Printf("- Selection: %s at odds %s\n", odd.Label, odd.Value)

					total := toInt(homeScore) + toInt(awayScore)
					var threshold float64
					var over bool

					if strings.HasPrefix(odd.Label, "Over") {
						fmt.Sscanf(odd.Label, "Over %f", &threshold)
						over = true
					} else {
						fmt.Sscanf(odd.Label, "Under %f", &threshold)
						over = false
					}

					if (over && float64(total) > threshold) || (!over && float64(total) < threshold) {
						fmt.Println("  ✅ Result: WON")
					} else {
						fmt.Println("  ❌ Result: LOST")
					}
				}
			}
		}
	}
}

func toInt(s string) int {
	var n int
	fmt.Sscanf(strings.TrimSpace(s), "%d", &n)
	return n
}

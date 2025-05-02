package models

// cricket
// CricketPrematchRoot represents the root of the cricket_prematch.json
type CricketPrematchRoot struct {
	Success int                 `json:"success"`
	Result  []CricketPrematch  `json:"result"`
}

type CricketPrematch struct {
	ID       string      `json:"id"`
	Time     string      `json:"time"`
	TimeStatus string    `json:"time_status"`
	League   League      `json:"league"`
	Home     Team        `json:"home"`
	Away     Team        `json:"away"`
	Odds     OddsWrapper `json:"odds"`
	Headers  []Header    `json:"headers"`
}



type OddsWrapper struct {
	PreMatch []MarketOdds `json:"pre_match"`
}

type MarketOdds struct {
	Label      string       `json:"label"`
	Bookmakers []Bookmaker  `json:"bookmakers"`
}

type Bookmaker struct {
	Name string            `json:"name"`
	Odds map[string]string `json:"odds"`
}



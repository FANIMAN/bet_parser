package parser

// parser/common


// League represents the details of a league
type League struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CC   string `json:"cc"`
}

// Team represents the details of a team
type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ImageID string `json:"image_id"`
	CC      string `json:"cc"`
}

// Header represents the headers in the prematch JSON
type Header struct {
	Name    string   `json:"name"`
	Markets []Market `json:"markets"`
}

// Market represents each market's name and odds in the prematch JSON
type Market struct {
	Name string `json:"name"`
	Odds []Odd  `json:"odds"`
}

// Odd represents the odds available for a market
type Odd struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

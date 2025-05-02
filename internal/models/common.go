package models


// common
type League struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CC   string `json:"cc"`
}

type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ImageID string `json:"image_id"`
	CC      string `json:"cc"`
}

type Header struct {
	Name    string   `json:"name"`
	Markets []Market `json:"markets"`
}

type Market struct {
	Name string `json:"name"`
	Odds []Odd  `json:"odds"`
}

type Odd struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

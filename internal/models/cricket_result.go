package models

// CricketResultRoot represents the root of cricket_result.json
type CricketResultRoot struct {
	Success int             `json:"success"`
	Results []CricketResult `json:"results"`
}

type CricketResult struct {
	ID        string         `json:"id"`
	SportID   string         `json:"sport_id"`
	Time      string         `json:"time"`
	TimeStatus string        `json:"time_status"`
	League    League         `json:"league"`
	Home      Team           `json:"home"`
	Away      Team           `json:"away"`
	Score     string         `json:"ss"` // ss = final score in format "home-away"
	Extra     ExtraData      `json:"extra"`
	ConfirmedAt string       `json:"confirmed_at"`
}

type ExtraData struct {
	StadiumData Stadium `json:"stadium_data"`
}

type Stadium struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Capacity     string `json:"capacity"`
	GoogleCoords string `json:"googlecoords"`
}

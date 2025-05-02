package models

// volleyball result
type VolleyballResultRoot struct {
	Success int                `json:"success"`
	Results []VolleyballResult `json:"results"`
}

type VolleyballResult struct {
	ID          string     `json:"id"`
	SportID     string     `json:"sport_id"`
	Time        string     `json:"time"`
	TimeStatus  string     `json:"time_status"`
	League      League     `json:"league"`
	Home        Team       `json:"home"`
	Away        Team       `json:"away"`
	Score       string     `json:"ss"`
	Extra       ExtraData  `json:"extra"`
	ConfirmedAt string     `json:"confirmed_at"`
}

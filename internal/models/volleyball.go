package models

type VolleyballPrematchRoot struct {
	Success int                  `json:"success"`
	Result  []VolleyballPrematch `json:"result"`
}

type VolleyballPrematch struct {
	ID         string      `json:"id"`
	Time       string      `json:"time"`
	TimeStatus string      `json:"time_status"`
	League     League      `json:"league"`
	Home       Team        `json:"home"`
	Away       Team        `json:"away"`
	Odds       OddsWrapper `json:"odds"`
	Headers    []Header    `json:"headers"`
}

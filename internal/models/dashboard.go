package models

type CategoryStat struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type SponsorStat struct {
	Sponsor string `json:"sponsor"`
	Count   int    `json:"count"`
}

type DashboardStats struct {
	TotalChildren int            `json:"total_children"`
	ByCategory    []CategoryStat `json:"by_category"`
	BySponsor     []SponsorStat  `json:"by_sponsor"`
}

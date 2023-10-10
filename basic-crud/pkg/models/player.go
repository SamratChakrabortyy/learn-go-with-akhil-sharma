package models

type Players struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	JersyNumber string `json:"jersyNumber"`
	Country     string `json:"country"`
	Team        *Team  `json:"team"`
}

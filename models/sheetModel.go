package models

type Sheet struct {
	Videourl string `json:"videourl"`
	Category string `json:"Category"`
	Name     string `json:"Name"`
	Link     string `json:"Link"`
	ID       int    `json:"Id"`
	Level    string `json:"Level"`
}

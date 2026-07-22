package models

type Quote struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

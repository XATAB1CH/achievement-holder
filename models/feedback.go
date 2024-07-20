package models

type Feedback struct {
	ID   uint   `json: "id"`
	Name string `json:  "name"`
	Text string `json: "text"`
}

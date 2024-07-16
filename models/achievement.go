package models

type Achievement struct {
	ID     uint   `json: "id"`
	Title  string `json:  "title"`
	Image  string `json:  "image" `
	Info   string `json: "info"`
	UserID uint   `jsom: "user_id"`
}

package achievement

type Achievement struct {
	ID     uint   `json: "id"`
	Title  string `json:  "title"`
	Img    string `json:  "img" `
	Info   string `json: "info"`
	UserID uint   `jsom: "user_id"`
}

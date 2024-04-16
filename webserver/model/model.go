package db_model

type FriendItem struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type ApplicationData struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

package models

type Post struct {
	Id          uint
	Title       string
	Description string
}

type PostPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

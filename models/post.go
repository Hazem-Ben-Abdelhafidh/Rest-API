package models

import (
	"encoding/json"
	"io"
)

type Post struct {
	Id          uint
	Title       string
	Description string
}

func (p *Post) Unmarshal(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Post) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)

}

type PostPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

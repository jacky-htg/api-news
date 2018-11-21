package models

import (
	"text/template"
	"time"
)

type Topic struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func TopicValidate(o Topic) Topic {
	o.Title = template.HTMLEscapeString(o.Title)
	return o
}

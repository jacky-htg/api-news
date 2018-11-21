package models

import (
	"database/sql"
	"text/template"
	"time"

	"github.com/go-sql-driver/mysql"
)

type News struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug,omitempty"`
	Content      string    `json:"content,omitempty"`
	Image        string    `json:"image,omitempty"`
	ImageCaption string    `json:"image_caption,omitempty"`
	Status       string    `json:"status,omitempty"`
	PublishDate  time.Time `json:"publish_date,omitempty"`
	Writer       User      `json:"writer,omitempty"`
	Editor       User      `json:"editor,omitempty"`
	Topic        []Topic   `json:"topic"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type NewsNull struct {
	Content      sql.NullString
	Image        sql.NullString
	ImageCaption sql.NullString
	PublishDate  mysql.NullTime
	Editor       sql.NullInt64
}

func NewsValidate(o News) News {
	o.Title = template.HTMLEscapeString(o.Title)
	o.Content = template.JSEscapeString(o.Content)

	return o
}

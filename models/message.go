package models

import (
	"time"
)

type MessageStruct struct {
	Title string    `json:"title"`
	Value string    `json:"value"`
	When  time.Time `json:"when"`
}

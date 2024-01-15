package model

import "time"

type Notification struct {
	ID        uint32    `json:"id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
}

package model

import "time"

type PermissionData map[string][]uint32

type Permission struct {
	ID        uint32         `json:"id"`
	UserID    uint32         `json:"user_id"`
	ObjectID  uint32         `json:"document_id"` // This might be removed
	Data      PermissionData `json:"permission"`
	CreatedAt time.Time      `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at" pg:"default:now()"`
}

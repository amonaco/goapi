package model

import "time"

type Document struct {
	ID          uint32                 `json:"id"`
	CompanyID   uint32                 `json:"company_id"`
	ProgramID   uint32                 `json:"program_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"name"`
	Data        map[string]interface{} `json:"data"`
	Status      map[string]interface{} `json:"status"`
	CreatedAt   time.Time              `json:"created_at" pg:"default:now()"`
	UpdatedAt   time.Time              `json:"updated_at" pg:"default:now()"`
}

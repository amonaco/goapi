package model

import "time"

type Quote struct {
	ID          uint32                 `json:"id"`
	UUID        string                 `json:"uuid"`
	UserID      uint32                 `json:"user_id"`
	CompanyID   uint32                 `json:"company_id"`
	ProgramID   uint32                 `json:"program_id"`
	FloorplanID uint32                 `json:"floorplan_id" pg:"floorplan_id,use_zero"`
	Title       string                 `json:"name"`
	Description string                 `json:"name"`
	Logo        string                 `json:"logo" pg:"-"`
	Number      uint32                 `json:"number"`
	Data        map[string]interface{} `json:"data"`
	Status      map[string]interface{} `json:"status"`
	CreatedAt   time.Time              `json:"created_at" pg:"default:now()"`
	UpdatedAt   time.Time              `json:"updated_at" pg:"default:now()"`
}

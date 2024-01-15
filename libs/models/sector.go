package model

import "time"

type Buyer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SectorData struct {
	Buyer Buyer `json:"buyer"`
}

type Sector struct {
	ID         uint32                 `json:"id"`
	UUID       string                 `json:"uuid"`
	ProgramID  uint32                 `json:"program_id"`
	CompanyID  uint32                 `json:"company_id"`
	UserID     uint32                 `json:"user_id"`
	Name       string                 `json:"name"`
	FloorNo    uint32                 `json:"floor_no" pg:":floor_no,use_zero"`
	Data       SectorData             `json:"data"`
	Floorplan  Floorplan              `json:"floorplan"`
	OriginalBF uint32                 `json:"original_bf_id" pg:"original_bf_id"`
	Status     map[string]interface{} `json:"status"`
	UpdatedAt  time.Time              `json:"updated_at" pg:"default:now()"`
}

type SectorWithOwner struct {
	Sector     `pg:",inherit"`
	OwnerEmail string `json:"owner_email"`
	OwnerName  string `json:"owner_name"`
}

type Floorplan struct {
	ID        uint32                 `json:"id"`
	SectorID  uint32                 `json:"sector_id"`
	CompanyID uint32                 `json:"company_id"`
	UserID    uint32                 `json:"user_id"`
	Name      string                 `json:"name"`    // Descriptive name for builder user
	Image     string                 `json:"image"`   // Image path of front-end generated image
	Version   uint32                 `json:"version"` // Incremental version for each Floorplan 1,2,3,4,5
	Data      map[string]interface{} `json:"data"`    // Actual floorplan JSON data
	Status    map[string]interface{} `json:"status"`
	UpdatedAt time.Time              `json:"updated_at" pg:"default:now()"`
	Quotes    []Quote                `json:"quotes"`
}

type BaseFloorplan struct {
	ID        uint32                 `json:"id"`
	SectorID  uint32                 `json:"sector_id"`
	Data      map[string]interface{} `json:"data"`
	Status    map[string]interface{} `json:"status"`
	UpdatedAt time.Time              `json:"updated_at" pg:"default:now()"`
}

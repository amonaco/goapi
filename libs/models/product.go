package model

import "time"

type Coordinates struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Group struct {
	Name string `json:"name"`
}

type Choice struct {
	Name      string `json:"name"`
	Group     uint32 `json:"group"`
	Thumbnail string `json:"thumbnail"`
	ID        string `json:"id"`
}

type Option struct {
	Name    string    `json:"name"`
	Groups  []*Group  `json:"groups,omitempty"`
	Choices []*Choice `json:"choices"`
	ID      string    `json:"id"`
}

type SizeOption struct {
	Step  uint32 `json:"step"`
	Start uint32 `json:"start"`
	End   uint32 `json:"end"`
}

type ProductData struct {
	Size        Coordinates `json:"size,omitempty"`
	Offset      Coordinates `json:"offset,omitempty"`
	Options     []*Option   `json:"options,omitempty"`
	Thumbnail   string      `json:"thumbnail"`
	DataSheet   string      `json:"data_sheet"`
	SizeOptions *SizeOption `json:"size_options"`
	// Next: fields used for materials
}

type Product struct {
	tableName   struct{}    `pg:",discard_unknown_columns"`
	ID          uint32      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Typ         string      `json:"typ"`
	CategoryID  uint32      `json:"-"`
	Category    *Category   `json:"category" db:"-"`
	BrandID     uint32      `json:"-"`
	Brand       *Brand      `json:"brand" db:"-"`
	Data        ProductData `json:"data"`
	Asset       *Asset      `json:"asset"`
	Thumbnail   string      `json:"thumbnail"`
	UUID        string      `json:"uuid"`
	CreatedAt   time.Time   `pg:"default:now()"`
	UpdatedAt   time.Time   `pg:"default:now()"`
}

type Brand struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID       uint32 `json:"id"`
	ParentID uint32 `json:"parent_id"`
	Name     string `json:"name"`
}

package model

type AssetData struct {
	Type   string                 `json:"type"`
	File   string                 `json:"file"`
	Size   map[string]interface{} `json:"size"`
	Offset map[string]interface{} `json:"offset"`
}

type Asset struct {
	ID         uint32    `json:"id"`
	ProductID  uint32    `json:"product_id"`
	Resolution string    `json:"resolution"`
	UUID       string    `json:"uuid"`
	Revision   int       `json:"revision"`
	Data       AssetData `json:"data" pg:"type:jsonb"`
}

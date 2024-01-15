package model

type Config struct {
	WorksDeadline     string `json:"works_deadline"`
	CoveringsDeadline string `json:"coverings_deadline"`
	WorksIntro        string `json:"works_intro"`
	WorksOutro        string `json:"works_outro"`
	CoveringsIntro    string `json:"coverings_intro"`
	CoveringsOutro    string `json:"coverings_outro"`
	BuilderIntro      string `json:"builder_intro"`
	BuilderOutro      string `json:"builder_outro"`
}

type ProgramData struct {
	Products []ProductPrice `json:"products"`
	Config   Config         `json:"config"`
	// Documents, etc.
}

type ProductPrice struct {
	ID    uint32  `json:"id"`
	Price float32 `json:"price"`
}

type PictureData struct {
	Pictures []Picture `json:"pictures"`
}

type Picture struct {
	Comment  string `json:"comment"`
	Filename string `json:"filename"`
}

type Program struct {
	ID         uint32      `json:"id"`
	CompanyID  uint32      `json:"company_id"`
	Name       string      `json:"name"`
	Address    string      `json:"address"`
	Thumbnail  string      `json:"thumbnail"`
	Flat_count uint32      `json:"flat_count"`
	Data       ProgramData `json:"data" pg:"type:jsonb"`
}

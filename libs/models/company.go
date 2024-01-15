package model

type Company struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
}

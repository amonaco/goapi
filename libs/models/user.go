package model

type User struct {
	ID        uint32   `json:"id" pg:"id,pk"`
	UUID      string   `json:"uuid" pg:"uuid"`
	CompanyID uint32   `json:"company_id" valid:"required"`
	Email     string   `json:"email" valid:"required,email"`
	Name      string   `json:"name" valid:"required"`
	Lastname  string   `json:"lastname" valid:"required"`
	Address   string   `json:"address" valid:"required"`
	Phone     string   `json:"phone"`
	Roles     []string `json:"roles"`
	Status    string   `json:"status"`
}

type Credentials struct {
	Token       string `json:"token" valid:"required"`
	Password    string `json:"password" valid:"required"`
	OldPassword string `json:"old_password"`
}

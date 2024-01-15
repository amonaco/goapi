package company

import (
	"context"
	"net/http"

	"github.com/go-pg/pg/v9"
	"github.com/jinzhu/copier"

	"github.com/amonaco/goapi/libs/auth"
	"github.com/amonaco/goapi/libs/database"
)

// put in model file
type CompanyModel struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
}

type CompanyRPC struct{}

// New creates a new CompanyRPC instance
func New() http.Handler {
	return NewCompanyServiceServer(&CompanyRPC{})
}

// Returns company details
func (a *CompanyRPC) GetCurrent(ctx context.Context) (*Company, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	company := &CompanyModel{}
	db := database.Get()
	err = db.Model(company).
		Where("company.id = ?", companyID).
		First()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	response := &Company{}
	copier.Copy(response, company)

	return response, nil
}

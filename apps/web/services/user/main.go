package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-pg/pg/v9"
	"github.com/jinzhu/copier"

	"github.com/amonaco/goapi/libs/auth"
	"github.com/amonaco/goapi/libs/cache"
	"github.com/amonaco/goapi/libs/config"
	"github.com/amonaco/goapi/libs/database"
	model "github.com/amonaco/goapi/libs/models"
	"github.com/amonaco/goapi/libs/worker"
)

// Implements the users service defined in rpc/users.ridl
type UserRPC struct{}

// New creates instanciates a new user service
func New() http.Handler {
	return NewUserServiceServer(&UserRPC{})
}

// Get returns a user by ID
func (u *UserRPC) Get(ctx context.Context, id uint32) (*User, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	user := &model.User{}
	db := database.Get()
	err = db.Model(user).
		Where("user.id = ?", id).
		Where("user.company_id = ?", companyID).
		First()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	response := &User{}
	copier.Copy(response, user)

	return response, nil
}

func (u *UserRPC) GetAll(ctx context.Context) ([]*User, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	db := database.Get()
	users := []*model.User{}
	err = db.Model(&users).
		Where("company_id = ?", companyID).
		Select()

	if err != nil {
		return nil, err
	}

	response := []*User{}
	copier.Copy(&response, &users)

	return response, nil
}

func createUser(ctx context.Context, u *User, role string) (uint32, error) {

	// Get company from session
	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return 0, err
	}

	user := &model.User{}
	copier.Copy(user, u)
	user.Roles = []string{role}
	user.CompanyID = companyID

	db := database.Get()
	_, err = db.Model(user).
		Returning("id").Insert()

	if err != nil {
		return 0, err
	}

	// Generate a random token for setting a password
	token, err := auth.GenerateToken()
	if err != nil {
		return 0, err
	}

	// Store the token in redis
	tokenKey := fmt.Sprintf("password_token:%d:%s", user.ID, token)
	err = cache.Set(tokenKey, user.Email, auth.TokenExpiry)
	if err != nil {
		return 0, err
	}

	// Build signin address
	conf := config.Get()
	tokenURL := fmt.Sprintf("%s/signin/%d:%s",
		conf.SigninURL, user.ID, token)

	// Send a message to the user
	w := worker.Get()
	task := w.NewTask("user_create")
	task.AddField("Name", user.Name)
	task.AddField("Email", user.Email)
	task.AddField("TokenURL", tokenURL)
	w.Push(task)

	return user.ID, nil
}

// Creates a editor user (console only)
func (u *UserRPC) CreateEditorUser(ctx context.Context, request *User) (*Status, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	// Get company from session
	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}
	request.CompanyID = companyID

	// Create the user
	userID, err := createUser(ctx, request, "editor")
	if err != nil {
		return nil, err
	}

	// Assign sector to user
	db := database.Get()
	sector := &model.Sector{}
	_, err = db.Model(sector).
		Set("user_id = ?", userID).
		Where("id = ?", request.SectorID).
		Where("company_id = ?", companyID).
		Update()

	if err != nil {
		return nil, err
	}

	response := &Status{
		Success: true,
		ID:      userID,
	}
	return response, err
}

func (u *UserRPC) Update(ctx context.Context, request *User) (*Status, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	// TODO: update RIDL to only accept restricted update values
	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: review ownership and isolation here
	user := &model.User{}
	copier.Copy(user, request)

	db := database.Get()
	_, err = db.Model(user).
		Where("id = ?", user.ID).
		Where("company_id = ?", companyID).
		Update()
	if err != nil {
		return nil, err
	}

	return &Status{Success: true}, nil
}

// TODO: Should disable a user, never delete
func (u *UserRPC) Disable(ctx context.Context, id uint32) (*Status, error) {
	return &Status{Success: true}, nil
}

// GetCurrent returns a user
func (u *UserRPC) GetCurrent(ctx context.Context) (*User, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user := &model.User{}
	db := database.Get()
	err = db.Model(user).
		Where("id = ?", userID).
		First()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	response := &User{}
	copier.Copy(response, user)

	return response, nil
}

func (p *UserRPC) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	// Get the companyID from session
	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	// Query with company_id
	asset := &model.User{}
	db := database.Get()
	err = db.Model(asset).
		Where("uuid = ?", uuid).
		Where("company_id = ?", companyID).
		First()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	response := &User{}
	copier.Copy(response, asset)

	return response, nil
}

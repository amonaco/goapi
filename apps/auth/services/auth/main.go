package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"

	"github.com/amonaco/goapi/libs/auth"
	"github.com/amonaco/goapi/libs/cache"
	"github.com/amonaco/goapi/libs/config"
	"github.com/amonaco/goapi/libs/database"
	model "github.com/amonaco/goapi/libs/models"
	"github.com/amonaco/goapi/libs/worker"
)

// This should probably go in ridl file
type User struct {
	ID        uint32   `json:"id"`
	CompanyID uint32   `json:"company_id"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Roles     []string `json:"roles"`
}

// Auth implements the auth service defined in rpc/auth.ridl
type AuthRPC struct{}

// New creates instanciates a new user service
func New() http.Handler {
	return NewAuthServiceServer(&AuthRPC{})
}

// Login creates a session for a user given the provided credentials are correct
func (a *AuthRPC) Login(ctx context.Context, creds *Credentials) (string, error) {
	if creds == nil {
		return "", Errorf(ErrInvalidArgument, "bad request")
	}

	_, err := govalidator.ValidateStruct(creds)
	if err != nil {
		return "", Errorf(ErrInvalidArgument, "invalid credentials")
	}

	user := User{}
	db := database.Get()
	err = db.Model(&user).
		Where("email = ?", creds.Email).
		Where("status->'confirmed' = 'true'").
		First()

	if err != nil {
		log.Println(err)
		return "", Errorf(ErrInvalidArgument, "wrong email or password")
	}

	hash := []byte(user.Password)
	pass := []byte(creds.Password)

	err = bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return "", Errorf(ErrInvalidArgument, "wrong email or password")
	}

	return login(ctx, user)
}

// Logout deletes the current session
func (a *AuthRPC) Logout(ctx context.Context, authToken string) (*Status, error) {

	err := logout(ctx, authToken)
	if err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func logout(ctx context.Context, token string) error {
	err := auth.DeleteSession(token)
	if err != nil {
		log.Println(err)
		return Errorf(ErrInternal, "an error has occurred")
	}

	w, ok := ctx.Value(HTTPResponseWriterCtxKey).(http.ResponseWriter)
	if !ok {
		log.Println("Critical Error! Request with no response writer!")
		return Errorf(ErrInternal, "an error has occurred")
	}

	cookie := http.Cookie{
		Name:   auth.TokenCookieName,
		Value:  "",
		MaxAge: -1, // MaxAge<0 means delete cookie now
	}
	http.SetCookie(w, &cookie)

	return nil
}

func login(ctx context.Context, user User) (string, error) {
	session, err := auth.CreateSession(user.ID, user.CompanyID, user.Roles)
	if err != nil {
		log.Println(err)
		return "", Errorf(ErrInternal, "an error has occurred")
	}

	w, ok := ctx.Value(HTTPResponseWriterCtxKey).(http.ResponseWriter)
	if !ok {
		log.Println("Critical Error! Request with no response writer!")
		return "", Errorf(ErrInternal, "an error has occurred")
	}

	expiry := time.Now().Add(7 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     auth.TokenCookieName,
		Value:    session.ID,
		Expires:  expiry,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return session.ID, nil
}

// SetPassword allows to change a user's own password
func (a *AuthRPC) SetPassword(ctx context.Context, request *PasswordToken) (*Status, error) {

	// Validate credentials model
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return nil, Errorf(ErrInvalidArgument, "invalid field: %s", err)
	}

	creds := &model.Credentials{}
	copier.Copy(creds, request)

	// Build token and check redis
	tokenKey := fmt.Sprintf("password_token:%s", creds.Token)
	email, err := cache.GetDel(tokenKey)
	if err != nil {
		return nil, err
	}

	// Validate email, could be malformed in redis
	valid := govalidator.IsEmail(email)
	if !valid {
		log.Printf("Set password request with invalid email")
		return nil, Errorf(ErrInvalidArgument, "invalid request")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Update user password
	user := &model.User{}
	db := database.Get()
	_, err = db.Model(user).
		Set("password = ?", string(hash)).
		Set("status = status || '{\"confirmed\": true}'").
		Where("email = ?", email).
		Update()
	if err != nil {
		return nil, err
	}

	return &Status{Success: true}, nil
}

func (a *AuthRPC) ForgotPassword(ctx context.Context, email string) error {

	// Validate email field
	valid := govalidator.IsEmail(email)
	if !valid {
		log.Printf("User password reset requested with bad email.")
		return nil
	}

	// Find the user e-mail
	user := &model.User{}
	db := database.Get()
	err := db.Model(user).
		Where("email = ?", email).
		First()

	// Return no response
	if err != nil {
		log.Printf("User password reset requested but not found on database.")
		return nil
	}

	// Generate a random token for setting a password
	token, err := auth.GenerateToken()
	if err != nil {
		return err
	}

	// Store the token in redis
	tokenKey := fmt.Sprintf("password_token:%s", token)
	err = cache.Set(tokenKey, user.Email, auth.TokenExpiry)
	if err != nil {
		return err
	}

	// Build signin address
	conf := config.Get()
	tokenURL := fmt.Sprintf("%s/signin/%s",
		conf.SigninURL, token)

	// Send a message to the user
	w := worker.Get()
	task := w.NewTask("user_forgot_password")
	task.AddField("Name", user.Name)
	task.AddField("Email", user.Email)
	task.AddField("TokenURL", tokenURL)
	w.Push(task)

	return nil
}

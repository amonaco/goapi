package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/amonaco/goapi/libs/database"
)

type PermissionData map[string][]uint32

type Permission struct {
	ID        uint32         `json:"id"`
	UserID    uint32         `json:"user_id"`
	ObjectID  uint32         `json:"document_id"` // This might be removed
	Data      PermissionData `json:"permission"`
	CreatedAt time.Time      `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at" pg:"default:now()"`
}

// Checks user_id permission to access an object
func Permissions(ctx context.Context, object string, id uint32) error {
	session, ok := getSessionFromContext(ctx)
	if !ok {
		log.Println("no session")
		return errors.New("auth: permission denied")
	}

	permission := &Permission{}
	db := database.Get()
	err := db.Model(permission).
		Where("permission.user_id = ?", session.UserID).
		Where("data @> '{\"program_id\": [?]}'", id).
		First()

	if err != nil {
		return errors.New("auth: permission rule format error")
	}

	return nil
}

// Authorize checks if a session has the required permissions
func Authorize(ctx context.Context, permissions ...string) error {
	if !isAuthorized(ctx, permissions) {
		return errors.New("auth: unauthorized")
	}

	return nil
}

func isAuthorized(ctx context.Context, permissions []string) bool {
	session, ok := getSessionFromContext(ctx)
	if !ok {
		log.Println("no session")
		return false
	}

	for _, permission := range permissions {
		for _, role := range session.Roles {
			if role == permission || role == "superadmin" {
				return true
			}
		}
	}

	return false
}

// GetUserID fetches the userID from the session
func GetUserID(ctx context.Context) (uint32, error) {
	session, ok := getSessionFromContext(ctx)
	if !ok {
		return 0, errors.New("auth: invalid user id")
	}

	return session.UserID, nil
}

// IsSuperAdmin checks if the current session is a superadmin
func IsSuperAdmin(ctx context.Context) bool {
	session, ok := getSessionFromContext(ctx)
	if !ok {
		return false
	}

	for _, role := range session.Roles {
		if role == "superadmin" {
			return true
		}
	}

	return false
}

// GetCompanyID extracts the companyID from the session
func GetCompanyID(ctx context.Context) (uint32, error) {
	session, ok := getSessionFromContext(ctx)
	if !ok {
		return 0, errors.New("auth: invalid company id")
	}

	return session.CompanyID, nil
}

func getSessionFromContext(ctx context.Context) (Session, bool) {
	session, ok := ctx.Value(ContextKey("session")).(Session)
	return session, ok
}

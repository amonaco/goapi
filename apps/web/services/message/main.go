package message

import (
	"context"
	"net/http"

	model "github.com/amonaco/goapi/libs/models"
	"github.com/go-pg/pg/v9"
	"github.com/jinzhu/copier"

	"github.com/amonaco/goapi/libs/auth"
	"github.com/amonaco/goapi/libs/database"
)

type MessageRPC struct{}

func New() http.Handler {
	return NewMessageServiceServer(&MessageRPC{})
}

func (u *MessageRPC) Get(ctx context.Context, id uint32) (*Message, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	message := &Message{}
	db := database.Get()
	err = db.Model(message).
		Where("message.id = ?", id).
		Where("message.company_id = ?", companyID).
		First()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return message, nil
}

func (u *MessageRPC) GetConversation(ctx context.Context, id uint32) ([]*Message, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	messages := []*Message{}
	db := database.Get()
	err = db.Model(messages).
		Where("message.subject = ?", id).
		Where("message.company_id = ?", companyID).
		Select()

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (u *MessageRPC) Create(ctx context.Context, request *Message) (*Status, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	message := &model.Message{}
	copier.Copy(message, request)
	message.CompanyID = companyID

	db := database.Get()
	_, err = db.Model(message).
		Returning("id").
		Insert()

	if err != nil {
		return nil, err
	}

	response := &Status{
		Success: true,
		ID:      message.ID,
	}
	return response, err
}

func (u *MessageRPC) Update(ctx context.Context, request *Message) (*Status, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	message := &model.Message{}
	copier.Copy(message, request)
	message.CompanyID = companyID

	db := database.Get()
	err = db.Update(message)
	if err != nil {
		return nil, err
	}

	response := &Status{Success: true}
	return response, err
}

func (u *MessageRPC) Remove(ctx context.Context, id uint32) (*Status, error) {
	err := auth.Authorize(ctx, "console")
	if err != nil {
		return nil, err
	}

	companyID, err := auth.GetCompanyID(ctx)
	if err != nil {
		return nil, err
	}

	message := &model.Message{}
	db := database.Get()
	_, err = db.Model(message).
		Where("message.id = ?", id).
		Where("message.company_id = ?", companyID).
		Delete()

	if err != nil {
		return nil, err
	}

	response := &Status{Success: true}
	return response, err
}

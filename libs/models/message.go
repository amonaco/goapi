package model

import (
	"time"
)

type Conversation struct {
	ID   uint32                 `json:"id"`
	Data map[string]interface{} `json:"data"`
	// Data can hold among other things
	// the IDs of the conversation members
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"default:now()"`
}

type Message struct {
	ID             uint32                 `json:"id"`
	UserID         uint32                 `json:"user_id"`
	CompanyID      uint32                 `json:"company_id"`
	ConversationID uint32                 `json:"conversation_id"`
	Data           map[string]interface{} `json:"data"`
	// Data holds members of the conversation
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"default:now()"`
}

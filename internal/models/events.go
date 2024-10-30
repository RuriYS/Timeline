package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseModel contains common fields for all collections
type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// MessageEvent represents a Discord message event
type MessageEvent struct {
	BaseModel `bson:",inline"`
	MessageID string                 `bson:"message_id"`
	GuildID   string                 `bson:"guild_id"`
	ChannelID string                 `bson:"channel_id"`
	UserID    string                 `bson:"user_id"`
	Content   string                 `bson:"content"`
	Metadata  map[string]interface{} `bson:"metadata"`
	Timestamp time.Time              `bson:"timestamp"`
}

func (m *MessageEvent) GetBase() *BaseModel {
	return &m.BaseModel
}

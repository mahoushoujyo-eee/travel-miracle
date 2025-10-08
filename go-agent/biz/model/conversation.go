package model

import (
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	UserId int64 `json:"user_id" gorm:"type:bigint;not null"`
	ConversationId string `json:"conversation_id" gorm:"type:varchar(255);not null"`
	Title string `json:"title" gorm:"type:varchar(255)"`
}

type ChatMemory struct {
	gorm.Model
	ConversationId string `json:"conversation_id" gorm:"type:varchar(255);not null"`
	Prompt string `json:"prompt" gorm:"type:text;not null"`
	Metadata string `json:"metadata" gorm:"type:varchar(255)"`
	Response string `json:"response" gorm:"type:text;not null"`
	Type string `json:"type" gorm:"type:varchar(255);not null"`
}

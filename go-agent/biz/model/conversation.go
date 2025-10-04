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
	Prompt string `json:"prompt" gorm:"type:varchar(255);not null"`
	ImgUrls []string `json:"img_urls" gorm:"type:varchar(255);not null"`
	Response string `json:"response" gorm:"type:varchar(255);not null"`
	Type string `json:"type" gorm:"type:varchar(255);not null"`
}

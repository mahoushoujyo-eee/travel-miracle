package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Score int `json:"score" gorm:"type:integer;not null"`
	Judgement string `json:"judgement" gorm:"type:text;"`
	ConversationID string `json:"conversation_id" gorm:"type:varchar(100);not null"`
}
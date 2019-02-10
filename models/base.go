package models

import "time"

// BaseModel contained in other models
type BaseModel struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"default:null"`
}

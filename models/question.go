package models

import (
	"github.com/jinzhu/gorm"
)

// Question is a
type Question struct {
	BaseModel
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}

// Create inserts the model into DB
func (question Question) Create(Db *gorm.DB) {
	Db.Create(&question)
}

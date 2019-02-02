package models

import (
	"github.com/jinzhu/gorm"
)

// Option is a
type Option struct {
	gorm.Model
	Label      string   `json:"label"`
	QuestionID uint     `json:"-"`
	Question   Question `json:"question"`
}

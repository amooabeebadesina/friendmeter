package models

import (
	"github.com/jinzhu/gorm"
)

// Question is a
type Question struct {
	gorm.Model
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}

var db := Db

func (question Question) create() {

}

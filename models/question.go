package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Question represents the question model
type Question struct {
	BaseModel
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}

// Create inserts the model into DB
func (question Question) Create(Db *gorm.DB) {
	Db.Create(&question)
}

// GetAll questions
func (question Question) GetAll(Db *gorm.DB) ([]Question, error) {
	var questions []Question

	err := Db.Find(&questions).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for i := range questions {
		Db.Model(&questions[i]).Related(&questions[i].Options)
	}
	return questions, nil
}

package models

// User is a struct of user
type User struct {
	BaseModel
	Name      string     `json:"name"`
	Questions []Question `gorm:"many2many:question_user;"`
}

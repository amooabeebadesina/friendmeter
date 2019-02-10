package models

// Option is a struct
type Option struct {
	BaseModel
	Label      string `json:"label"`
	QuestionID uint   `json:"-" sql:"index"`
}

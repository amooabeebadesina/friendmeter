package controllers

import (
	"encoding/json"
	"html/template"
	"io"
	"knowme/models"
	"knowme/utils"
)

type Home struct {
}

// Index function resolves the / route
func (home Home) Index() (temp *template.Template) {
	paths := []string{
		"templates/index.html",
	}

	t, _ := template.ParseFiles(paths...)
	return t
}

// CreateQuestions creates a question instance
func (home Home) CreateQuestions(Db *gorm.Db, data io.ReadCloser) []byte {
	var question models.Question
	if err := json.NewDecoder(data).Decode(&question); err != nil {
		panic(err)
	}

	resData := utils.JSONResponse{
		Status: "success",
	}

	res, _ := json.Marshal(resData)

	return res

}

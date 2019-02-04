package main

import (
	"encoding/json"
	"html/template"
	"knowme/models"
	"knowme/utils"
	"net/http"

	"github.com/go-chi/chi"
)

var jsonResponse utils.JSONResponse

func indexRoute(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"templates/index.html",
	}
	t, _ := template.ParseFiles(paths...)
	t.Execute(w, nil)
}

func addQuestions(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	authorization := r.Header.Get("authorization")
	if authorization == "" {
		res, _ := json.Marshal("You might have lost your way")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		return
	}

	data := r.Body

	if err := json.NewDecoder(data).Decode(&question); err != nil {
		panic(err)
	}

	question.Create(Db)

	utils.SendSuccessResponse(w, question)

}

func routers() *chi.Mux {

	router.Get("/", indexRoute)
	router.Post("/superhero", addQuestions)

	return router
}

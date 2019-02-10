package main

import (
	"encoding/json"
	"html/template"
	"knowme/models"
	"knowme/utils"
	"net/http"

	"github.com/go-chi/chi"
)

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

func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	res, err := question.GetAll(Db)
	if err != nil {
		utils.SendErrorResponse(w, "Error fetching questions")
	} else {
		utils.SendSuccessResponse(w, res)
	}
}

func routers() *chi.Mux {

	router.Get("/", indexRoute)
	router.Post("/superhero", addQuestions)

	router.Route("/api", func(r chi.Router) {
		r.Get("/questions", getAllQuestions)
	})

	return router
}

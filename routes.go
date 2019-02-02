package main

import (
	"encoding/json"
	"knowme/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

var homeController = controllers.Home{}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	t := homeController.Index()
	t.Execute(w, nil)
}

func addQuestions(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("authorization")
	if authorization == "" {
		res, _ := json.Marshal("Get the fuck out nigga")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		return
	}
	data := r.Body
	res := homeController.CreateQuestions(&Db, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func routers() *chi.Mux {

	router.Get("/", indexRoute)
	router.Post("/superhero", addQuestions)

	return router
}

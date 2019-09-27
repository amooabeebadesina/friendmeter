package main

import (
	"models"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var router *chi.Mux
var Db *gorm.DB
var tokenAuth *jwtauth.JWTAuth

func initDatabase() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	Db, err = gorm.Open(
		"postgres",
		"sslmode=disable host="+os.Getenv("DATABASE_HOST")+
			" user="+os.Getenv("DATABASE_USERNAME")+
			" password="+os.Getenv("DATABASE_PASSWORD")+
			" dbname="+os.Getenv("DATABASE_NAME"))
	if err != nil {
		panic(err)
	}
}

func initMiddlewares() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(cors.Handler)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}

func init() {
	initDatabase()
	Db.AutoMigrate(&models.User{}, &models.Question{}, &models.Option{})
	router = chi.NewRouter()
	router = routers()
}

func main() {
	defer Db.Close()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

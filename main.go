package main

import (
	"go-chi-gorm/router"
	"log"
	"net/http"

	_ "go-chi-gorm/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title go-chi-gorm
// @version 0.1
// @description This is a practice.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	router.NewUserRouter(r, db).Init()
	router.NewPostRouter(r, db).Init()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

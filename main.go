package main

import (
	"go-mux-gorm/router"
	"log"
	"net/http"

	_ "go-mux-gorm/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title go-mux-gorm
// @version 0.1
// @description This is a practice.
func main() {
	r := mux.NewRouter()

	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	router.NewUserRouter(r, db).Init()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.Use(CORSMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

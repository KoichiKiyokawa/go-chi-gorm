package router

import (
	"go-mux-gorm/handler"
	"go-mux-gorm/repository"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserRouter struct {
	router  *mux.Router
	handler *handler.UserHandler
	db      *gorm.DB
}

func NewUserRouter(r *mux.Router, db *gorm.DB) *UserRouter {
	userRepo := repository.NewUserRepository(db)
	return &UserRouter{router: r, db: db, handler: handler.NewUserHandler(userRepo)}
}

func (u *UserRouter) Init() {
	u.router.HandleFunc("/users", u.handler.Index).Methods(http.MethodGet)
	u.router.HandleFunc("/users/{id}", u.handler.Show).Methods(http.MethodGet)
	u.router.HandleFunc("/users", u.handler.Create).Methods(http.MethodPost)
}

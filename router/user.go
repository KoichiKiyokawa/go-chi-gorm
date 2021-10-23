package router

import (
	"go-chi-gorm/handler"
	"go-chi-gorm/repository"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type UserRouter struct {
	router  *chi.Mux
	handler *handler.UserHandler
	db      *gorm.DB
}

func NewUserRouter(r *chi.Mux, db *gorm.DB) *UserRouter {
	userRepo := repository.NewUserRepository(db)
	return &UserRouter{router: r, db: db, handler: handler.NewUserHandler(userRepo)}
}

func (r *UserRouter) Init() {
	r.router.Route("/users", func(s chi.Router) {

		s.Get("/", r.handler.Index)
		s.Get("/{id}", r.handler.Show)
		s.Post("/", r.handler.Create)
	})
}

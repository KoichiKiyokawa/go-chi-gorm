package router

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

type CoreRouter struct {
	router *chi.Mux
	db     *gorm.DB
}

func NewCoreRouter(router *chi.Mux, db *gorm.DB) *CoreRouter {
	return &CoreRouter{router: router, db: db}
}

func (r *CoreRouter) Init() {
	newUserRouter(r.router, r.db).Init()
	newPostRouter(r.router, r.db).Init()

	r.router.Get("/swagger/*", httpSwagger.WrapHandler)
}

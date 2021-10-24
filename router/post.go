package router

import (
	"go-chi-gorm/handler"
	"go-chi-gorm/repository"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type PostRouter struct {
	router  *chi.Mux
	handler *handler.PostHandler
	db      *gorm.DB
}

func newPostRouter(r *chi.Mux, db *gorm.DB) *PostRouter {
	postRepo := repository.NewPostRepository(db)
	return &PostRouter{router: r, db: db, handler: handler.NewPostHandler(postRepo)}
}

func (r *PostRouter) Init() {
	r.router.Route("/posts", func(s chi.Router) {

		s.Get("/", r.handler.Index)
		s.Get("/{id}", r.handler.Show)
		s.Post("/", r.handler.Create)
	})

	r.router.Get("/users/{id}/posts", r.handler.UserPostsIndex)
}

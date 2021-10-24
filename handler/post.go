package handler

import (
	"encoding/json"
	"go-chi-gorm/entity"
	"go-chi-gorm/repository"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type PostHandler struct {
	postRepo *repository.PostRepository
}

func NewPostHandler(postRepo *repository.PostRepository) *PostHandler {
	return &PostHandler{postRepo: postRepo}
}

// @Summary 投稿一覧
// @Description 投稿一覧を返却する
// @Tags post
// @Success 200 {object} []entity.Post
// @Router /posts [get]
func (h *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := h.postRepo.FindAll()
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, users)
}

// @Summary 投稿詳細
// @Description 指定した投稿を返却する
// @Tags post
// @Param id path string true "投稿のID"
// @Success 200 {object} entity.Post
// @Router /posts/{id} [get]
func (h *PostHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, err)
		return
	}
	user, err := h.postRepo.FindOne(id)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, user)
}

// @Summary ユーザの投稿一覧
// @Description あるユーザの投稿一覧を返却する
// @Tags user
// @Success 200 {object} []entity.Post
// @Router /users/{id}/posts [get]
func (h *PostHandler) UserPostsIndex(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, err)
		return
	}
	posts, err := h.postRepo.FindAllByUserId(userID)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, posts)
}

// @Summary 投稿作成
// @Description 投稿を作成する
// @Tags post
// @Param post_data body entity.Post true "作成する投稿のデータ"
// @Success 200 {object} entity.Post
// @Router /posts [post]
func (u *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, user)
}

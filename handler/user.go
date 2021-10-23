package handler

import (
	"encoding/json"
	"go-mux-gorm/entity"
	"go-mux-gorm/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// @Summary ユーザ一覧
// @Description ユーザ一覧を返却する
// @Success 200 {object} []entity.User
// @Router /users/{id} [get]
func (u *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, users)
}

func (u *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, err)
		return
	}
	user, err := u.userRepo.FindOne(id)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, user)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, user)
}

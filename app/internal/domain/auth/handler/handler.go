package handler

import (
	"encoding/json"
	"io"
	"net/http"

	apperror "github.com/kurochkinivan/Note-taker/internal/appError"
	"github.com/kurochkinivan/Note-taker/internal/domain/auth/model"
	"github.com/kurochkinivan/Note-taker/internal/domain/auth/repository"
	"github.com/kurochkinivan/Note-taker/internal/domain/handlers"
)

type handler struct {
	repository repository.Repository
}

func NewAuthHandler(repository repository.Repository) handlers.Handler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) Register(mux *http.ServeMux) {
	mux.HandleFunc(http.MethodGet+" /auth/sign-in", apperror.MiddleWare(h.signIn))
}

func (h *handler) signIn(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return apperror.ErrValidateData
	}
	defer r.Body.Close()

	var user model.User
	if err := json.Unmarshal(data, &user); err != nil {
		return apperror.ErrSerializeData
	}

	user, err = h.repository.GetUser(user.Login, user.Password)
	if err != nil {
		return err
	}

	jwt, err := h.repository.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	jwtResponse := model.JWTResponse{
		JWT: jwt,
	}

	resp, err := json.Marshal(jwtResponse)
	if err != nil {
		return apperror.ErrSerializeData
	}

	w.Write(resp)

	return nil
}

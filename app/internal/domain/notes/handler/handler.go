package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	apperror "github.com/kurochkinivan/Note-taker/internal/appError"
	"github.com/kurochkinivan/Note-taker/internal/domain/handlers"
	"github.com/kurochkinivan/Note-taker/internal/domain/notes/model"
	"github.com/kurochkinivan/Note-taker/internal/domain/notes/repository"
	yaspeller "github.com/kurochkinivan/Note-taker/internal/external/yaSpeller"
	mdw "github.com/kurochkinivan/Note-taker/internal/middleware"
)

type handler struct {
	repository repository.Repository
}

func NewNotesRepository(repository repository.Repository) handlers.Handler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) Register(mux *http.ServeMux) {
	mux.HandleFunc(http.MethodGet+" /notes/all", mdw.ErrorMiddleware(mdw.AuthMiddleware(h.GetAll)))
	mux.HandleFunc(http.MethodPost+" /notes/create", mdw.ErrorMiddleware(mdw.AuthMiddleware(h.CreateNote)))
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Header.Get("user_id")

	notes, err := h.repository.GetAll(context.TODO(), userID)
	if err != nil {
		return err
	}

	data, err := json.Marshal(notes)
	if err != nil {
		return apperror.ErrSerializeData
	}

	w.Write(data)

	return nil
}

func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return apperror.ErrValidateData
	}
	defer r.Body.Close()

	var note model.Note
	err = json.Unmarshal(data, &note); 
	if err != nil {
		return apperror.ErrSerializeData
	}

	note.Body, err = yaspeller.CorrectMistakes(note.Body)
	if err != nil {
		return err
	}

	note.UserID = r.Header.Get("user_id")
	err = h.repository.Create(context.TODO(), note)
	if err != nil {
		return err
	}

	return nil
}

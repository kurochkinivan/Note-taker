package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func MiddleWare(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)

		if err != nil {
			// Errors are sorted by the status codes
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrValidateData) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(ErrNotFound.Marshal())
					return
				}

				if errors.Is(err, ErrSerializeData) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(ErrSerializeData.Marshal())
					return
				}

				if errors.Is(err, ErrInvalidPassword) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(ErrInvalidPassword.Marshal())
					return
				}

				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				if errors.Is(err, ErrSignToken) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(ErrSerializeData.Marshal())
					return
				}
			}
			
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(systemError(err).Marshal())
		}
	}
}

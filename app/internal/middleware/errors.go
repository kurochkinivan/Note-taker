package middleware

import (
	"errors"
	"net/http"

	aerr "github.com/kurochkinivan/Note-taker/internal/appError"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorMiddleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *aerr.AppError
		err := h(w, r)

		// Можно сделать это гораздо короче, если включить в ошибку http статус-код.
		//
		// В данном проекте я решил попробовать применить такой подход, чтобы изолировать
		// транспортный протокол от кастомной ошибки. (если, например, нужна будет миграция на gRPC)
		if err != nil {
			// Errors are sorted by the status codes
			if errors.As(err, &appErr) {
				// ErrValidateData - 400
				if errors.Is(err, aerr.ErrValidateData) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(aerr.ErrNotFound.Marshal())
					return
				}

				// ErrInvalidSigningMethod - 400
				if errors.Is(err, aerr.ErrInvalidSigningMethod) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(aerr.ErrInvalidSigningMethod.Marshal())
					return
				}

				
				// ErrEmptyAuthHeader - 400
				if errors.Is(err, aerr.ErrEmptyAuthHeader) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(aerr.ErrEmptyAuthHeader.Marshal())
					return
				}
				
				// ErrInvalidAuthHeader - 400
				if errors.Is(err, aerr.ErrInvalidAuthHeader) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(aerr.ErrInvalidAuthHeader.Marshal())
					return
				}
				
				// ErrSerializeData - 400
				if errors.Is(err, aerr.ErrSerializeData) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(aerr.ErrSerializeData.Marshal())
					return
				}
				
				// ErrInvalidPassword - 401
				if errors.Is(err, aerr.ErrInvalidPassword) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(aerr.ErrInvalidPassword.Marshal())
					return
				}
				
				// ErrTokenExired - 401
				if errors.Is(err, aerr.ErrTokenExired) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(aerr.ErrTokenExired.Marshal())
					return
				}

				// ErrNotFound - 404
				if errors.Is(err, aerr.ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(aerr.ErrNotFound.Marshal())
					return
				}

				// ErrSignToken - 500
				if errors.Is(err, aerr.ErrSignToken) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(aerr.ErrSerializeData.Marshal())
					return
				}

				// ErrAssertingJWT - 500
				if errors.Is(err, aerr.ErrAssertingJWT) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(aerr.ErrAssertingJWT.Marshal())
					return
				}
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(aerr.SystemError(err).Marshal())
		}
	}
}

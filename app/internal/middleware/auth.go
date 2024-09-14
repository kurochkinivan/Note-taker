package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	apperror "github.com/kurochkinivan/Note-taker/internal/appError"
	"github.com/kurochkinivan/Note-taker/internal/constants"
)

func AuthMiddleware(h appHandler) appHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		header := r.Header.Get("Authorization")
		if header == "" {
			return apperror.ErrEmptyAuthHeader
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return apperror.ErrInvalidAuthHeader
		}

		payload, err := parseToken(headerParts[1])
		if err != nil {
			return err
		}

		expTime := time.Unix(int64(payload["exp"].(float64)), 0)
		if time.Now().After(expTime) {
			return apperror.ErrTokenExired
		}

		r.Header.Set("user_id", payload["ueid"].(string))

		err = h(w, r)
		if err != nil {
			return err
		}

		return nil
	}
}

func parseToken(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperror.ErrInvalidSigningMethod
		}
		return []byte(constants.Signingkey), nil
	})
	if err != nil {
		return nil, apperror.ErrInvalidAuthHeader
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apperror.ErrAssertingJWT
	}

	return payload, nil
}

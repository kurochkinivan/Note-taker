package repository

import "github.com/kurochkinivan/Note-taker/internal/domain/auth/model"

type Repository interface {
	GetUser(login, password string) (model.User, error)

	GenerateToken(uuid string) (string, error)
}

package repository

import (
	"context"

	"github.com/kurochkinivan/Note-taker/internal/domain/notes/model"
)

type Repository interface {
	Create(ctx context.Context, note model.Note) error
	GetAll(ctx context.Context, userID string) ([]model.Note, error)
}

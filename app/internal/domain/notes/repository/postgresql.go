package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kurochkinivan/Note-taker/internal/constants"
	"github.com/kurochkinivan/Note-taker/internal/domain/notes/model"
	psql "github.com/kurochkinivan/Note-taker/pkg/postgresql"
	"github.com/sirupsen/logrus"
)

type notesRepository struct {
	client psql.PostgreSQLClient
	qb     sq.StatementBuilderType
}

func NewNotesRepository(client psql.PostgreSQLClient) Repository {
	return &notesRepository{
		client: client,
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repo *notesRepository) Create(ctx context.Context, note model.Note) error {
	logrus.Tracef("creating new note for user with id %q", note.UserID)

	sql, args, err := repo.qb.
		Insert(constants.NotesTable).
		Columns(
			"user_id",
			"title",
			"body",
		).
		Values(
			&note.UserID,
			&note.Title,
			&note.Body,
		).
		ToSql()
	if err != nil {
		return psql.ErrCreateQuery(err)
	}

	cmd, err := repo.client.Exec(ctx, sql, args...)
	if err != nil {
		return psql.ErrExec(err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("nothing interested")
	}

	logrus.Tracef("rows affected: %d", cmd.RowsAffected())

	return nil
}

func (repo *notesRepository) GetAll(ctx context.Context, userID string) ([]model.Note, error) {
	logrus.Tracef("getting all notes for user with id: %q", userID)

	sql, args, err := repo.qb.
		Select(
			"title",
			"body",
			"created_at",
		).
		From(constants.NotesTable).
		Where(sq.Eq{"user_id": userID}). 
		ToSql()
	if err != nil {
		return nil, psql.ErrCreateQuery(err)
	}

	rows, err := repo.client.Query(ctx, sql, args...)
	if err != nil {
		return nil, psql.ErrDoQuery(err)
	}

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		
		if err := rows.Scan(
			&note.Title,
			&note.Body,
			&note.CreatedAt,
		); err != nil {
			return nil, psql.ErrScan(err)
		}

		notes = append(notes, note)
	}

	return notes, nil
}

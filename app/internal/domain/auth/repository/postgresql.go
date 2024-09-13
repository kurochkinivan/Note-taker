package repository

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/kurochkinivan/Note-taker/internal/appError"
	"github.com/kurochkinivan/Note-taker/internal/constants"
	"github.com/kurochkinivan/Note-taker/internal/domain/auth/model"
	psql "github.com/kurochkinivan/Note-taker/pkg/postgresql"
	"github.com/sirupsen/logrus"
)

type authRepository struct {
	qb     sq.StatementBuilderType
	client psql.PostgreSQLClient
}

func NewAuthRepository(client psql.PostgreSQLClient) Repository {
	return &authRepository{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *authRepository) GetUser(login, password string) (model.User, error) {
	logrus.Tracef("getting user with login: %q", login)

	sql, args, err := repo.qb.
		Select(
			"id",
			"login",
			"password",
		).
		From(constants.UsersTable).
		Where(sq.And{
			sq.Eq{"login": login},
		}).
		ToSql()
	if err != nil {
		return model.User{}, psql.ErrCreateQuery(err)
	}

	var user model.User
	err = repo.client.QueryRow(context.TODO(), sql, args...).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, apperror.ErrNotFound
		} 
		return model.User{}, psql.ErrScan(err)
	}

	if user.Password != hashPassword(password) {
		return model.User{}, apperror.ErrInvalidPassword
	}

	return user, nil
}

func (repo *authRepository) GenerateToken(uuid string) (string, error) {
	logrus.Tracef("generating token for user with uuid %q", uuid)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ueid": uuid,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(constants.TokenTTL).Unix(),
	})

	signedToken, err := token.SignedString([]byte(constants.Signingkey))
	if err != nil {
		return "", apperror.ErrSignToken
	}

	return signedToken, nil
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(constants.Salt)))
}

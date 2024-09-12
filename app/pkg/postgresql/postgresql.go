package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgreSQLClient interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

func NewClient(ctx context.Context, maxAttempts int, pgConfig *PgConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", pgConfig.userName, pgConfig.password, pgConfig.host, pgConfig.port, pgConfig.database)

	conn, parseErr := pgxpool.ParseConfig(connString)
	if parseErr != nil {
		logrus.Errorf("unable to parse psql config due to err: %v\n", parseErr)
		return nil, parseErr
	}

	pool, connErr := pgxpool.NewWithConfig(ctx, conn)
	if connErr != nil {
		logrus.Errorf("failed to create pgxpool due to err: %v\n", connErr)
		return nil, connErr
	}
	
	err := doWithAttempts(func() error {
		pingErr := pool.Ping(ctx)
		if pingErr != nil {
			logrus.Infof("Failed to connect to postgres, err: %v... Going to do next attempt\n", pingErr)
			return pingErr
		}

		return nil
	}, maxAttempts, 5 * time.Second)
	if err != nil {
		logrus.Fatalf("All attempts are exceeded. Unable to connect to PostgreSQL, err: %v", err)
		return nil, err
	}

	return pool, nil
}
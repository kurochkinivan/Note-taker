package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/kurochkinivan/Note-taker/internal/config"
	authHandler "github.com/kurochkinivan/Note-taker/internal/domain/auth/handler"
	authRepository "github.com/kurochkinivan/Note-taker/internal/domain/auth/repository"
	notesHandler "github.com/kurochkinivan/Note-taker/internal/domain/notes/handler"
	notesRepository "github.com/kurochkinivan/Note-taker/internal/domain/notes/repository"
	"github.com/kurochkinivan/Note-taker/pkg/postgresql"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type App struct {
	cfg        *config.Config
	mux        *http.ServeMux
	httpserver *http.Server
}

func NewApp(cfg *config.Config) (*App, error) {
	logrus.Info("initializing router")
	mux := http.NewServeMux()

	logrus.Info("postgreSQL initializing")
	cfgPSQL := cfg.PostgreSQL
	logrus.WithFields(logrus.Fields{
		"username": cfgPSQL.Username,
		"password": cfgPSQL.Password,
		"host":     cfgPSQL.Host,
		"port":     cfgPSQL.Port,
		"database": cfgPSQL.Database,
	}).Debug("postgresql credentials")

	pgcfg := postgresql.NewPgConfig(cfgPSQL.Username, cfgPSQL.Password, cfgPSQL.Host, cfgPSQL.Port, cfgPSQL.Database)
	client, err := postgresql.NewClient(context.TODO(), 5, pgcfg)
	if err != nil {
		return nil, fmt.Errorf("failed to conect to postgresql database due to error: %q", err)
	}

	authRepository := authRepository.NewAuthRepository(client)
	authHandler := authHandler.NewAuthHandler(authRepository)
	authHandler.Register(mux)

	notesRepository := notesRepository.NewNotesRepository(client)
	notesHandler := notesHandler.NewNotesRepository(notesRepository)
	notesHandler.Register(mux)

	return &App{
		cfg: cfg,
		mux: mux,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	errg, _ := errgroup.WithContext(ctx)
	errg.Go(func() error {
		return a.startHTTP()
	})
	logrus.Info("application is initialized and stated")

	return errg.Wait()
}

func (a *App) startHTTP() error {
	logrus.Info("creating listener")
	logrus.WithFields(logrus.Fields{
		"IP":   a.cfg.HTTP.IP,
		"Port": a.cfg.HTTP.Port,
	}).Debug("http listener credentials")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		logrus.WithError(err).Error("failed to create listener")
		return err
	}

	a.httpserver = &http.Server{
		Handler:      a.mux,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	logrus.Info("application completely initialized and started")

	if err = a.httpserver.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logrus.Warn("server shutdown")
		default:
			logrus.Fatal(err)
		}
	}

	err = a.httpserver.Shutdown(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}

	return err
}

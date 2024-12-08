package app

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"sso/pkg/dbconnector/postgresql"
)

type Application struct {
	cfg Config
	db  Database
	log *slog.Logger
}

func New(cfg Config, log *slog.Logger) *Application {
	return &Application{
		cfg: cfg,
		log: log,
	}
}

func (a *Application) Start() error {
	databaseConfig := a.cfg.Database()
	switch databaseConfig.Database() {
	case postgresql.DatabaseName:
		database, err := postgresql.New(databaseConfig)
		if err != nil {
			return err
		}
		a.db = database
	default:
		return errors.New("database not supported")
	}
	err := a.db.Migrate()
	if err != nil {
		return fmt.Errorf("migrate failed: %w", err)
	}

	return nil
}

func (a *Application) Stop() {
	err := a.db.Close()
	if err != nil {
		a.log.Error("failed to close database", slog.String("error", err.Error()))
	}
}

type Database interface {
	Connection() *sqlx.DB
	Migrate() error
	Close() error
}

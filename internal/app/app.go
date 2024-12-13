package app

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"sso/internal/api/grpc"
	"sso/internal/api/grpc/controller"
	"sso/internal/errs"
	"sso/internal/infrastructure/keystorage"
	"sso/internal/usecase/keyuc"
	"sso/pkg/dbconnector/postgresql"
)

type Application struct {
	cfg        Config
	db         Database
	grpcServer *grpc.Server
	log        *slog.Logger
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

	keyStorage := keystorage.New(a.log)
	err = keyStorage.SetRefreshSecret(a.cfg.RefreshToken())
	if err != nil {
		return errs.Wrap("failed to set refresh secret", err)
	}
	err = keyStorage.SetAccessSecret(a.cfg.AccessToken())
	if err != nil {
		return errs.Wrap("failed to set access secret", err)
	}

	keyUseCase := keyuc.New(keyStorage, a.log)

	grpcController := controller.New(keyUseCase, nil, nil, a.log)

	a.grpcServer = grpc.New(a.cfg.GRPC(), grpcController, a.log)
	err = a.grpcServer.Start()
	if err != nil {
		return errs.Wrap("failed to start grpc server", err)
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

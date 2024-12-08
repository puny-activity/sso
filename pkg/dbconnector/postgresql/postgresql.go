package postgresql

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"sso/pkg/dbconnector"
)

const DatabaseName = "postgresql"

type Postgresql struct {
	db            *sqlx.DB
	migrationPath string
}

func New(cfg dbconnector.Config) (*Postgresql, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable default_query_exec_mode=cache_describe",
		cfg.Host(), cfg.Port(), cfg.User(), cfg.Name(), cfg.Password())

	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgresql connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgresql connection: %w", err)
	}

	return &Postgresql{
		db:            db,
		migrationPath: cfg.MigrationsPath(),
	}, nil
}

func (p *Postgresql) Connection() *sqlx.DB {
	return p.db
}

func (p *Postgresql) Migrate() error {
	goose.SetTableName("migrations")

	err := goose.Up(p.db.DB, p.migrationPath)
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
func (p *Postgresql) Close() error {
	err := p.db.Close()
	if err != nil {
		return fmt.Errorf("failed to close postgresql connection: %w", err)
	}
	return nil
}

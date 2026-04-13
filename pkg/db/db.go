package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tecknohow/TaskDream/pkg/config"
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"
)

var Engine *xorm.Engine

func InitDB(cfg *config.Config) (*xorm.Engine, error) {
	var engine *xorm.Engine
	var err error

	switch cfg.Database.Driver {
	case "sqlite":
		engine, err = xorm.NewEngine("sqlite3", cfg.Database.SQLite.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite engine: %w", err)
		}

	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Database.Postgres.Host,
			cfg.Database.Postgres.Port,
			cfg.Database.Postgres.User,
			cfg.Database.Postgres.Password,
			cfg.Database.Postgres.Database,
			cfg.Database.Postgres.SSLMode,
		)
		engine, err = xorm.NewEngine("postgres", dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to create postgres engine: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	// Set logger
	engine.SetLogLevel(xormlog.LOG_DEBUG)

	// Test connection
	if err := engine.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return engine, nil
}

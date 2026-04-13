package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type DatabaseConfig struct {
	Driver   string
	SQLite   SQLiteConfig
	Postgres PostgresConfig
}

type SQLiteConfig struct {
	Path string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/taskdream/")
	viper.AddConfigPath(os.ExpandEnv("$HOME/.taskdream"))

	// Set defaults
	viper.SetDefault("server.port", 3456)
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.sqlite.path", "./taskdream.db")
	viper.SetDefault("database.postgres.host", "localhost")
	viper.SetDefault("database.postgres.port", 5432)
	viper.SetDefault("database.postgres.user", "taskdream")
	viper.SetDefault("database.postgres.database", "taskdream")
	viper.SetDefault("database.postgres.sslmode", "disable")
	viper.SetDefault("jwt.secret", "your-secret-key-change-in-production")

	// Read environment variables
	viper.BindEnv("server.port", "TASKDREAM_SERVER_PORT")
	viper.BindEnv("server.host", "TASKDREAM_SERVER_HOST")
	viper.BindEnv("database.driver", "TASKDREAM_DATABASE_DRIVER")
	viper.BindEnv("database.sqlite.path", "TASKDREAM_SQLITE_PATH")
	viper.BindEnv("database.postgres.host", "TASKDREAM_POSTGRES_HOST")
	viper.BindEnv("database.postgres.port", "TASKDREAM_POSTGRES_PORT")
	viper.BindEnv("database.postgres.user", "TASKDREAM_POSTGRES_USER")
	viper.BindEnv("database.postgres.password", "TASKDREAM_POSTGRES_PASSWORD")
	viper.BindEnv("database.postgres.database", "TASKDREAM_POSTGRES_DATABASE")
	viper.BindEnv("jwt.secret", "TASKDREAM_JWT_SECRET")

	// Try to read config file, but don't fail if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: viper.GetInt("server.port"),
			Host: viper.GetString("server.host"),
		},
		Database: DatabaseConfig{
			Driver: viper.GetString("database.driver"),
			SQLite: SQLiteConfig{
				Path: viper.GetString("database.sqlite.path"),
			},
			Postgres: PostgresConfig{
				Host:     viper.GetString("database.postgres.host"),
				Port:     viper.GetInt("database.postgres.port"),
				User:     viper.GetString("database.postgres.user"),
				Password: viper.GetString("database.postgres.password"),
				Database: viper.GetString("database.postgres.database"),
				SSLMode:  viper.GetString("database.postgres.sslmode"),
			},
		},
		JWT: JWTConfig{
			Secret: viper.GetString("jwt.secret"),
		},
	}

	return cfg, nil
}

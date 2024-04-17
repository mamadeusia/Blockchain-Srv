package config

import "fmt"

type Postgres struct {
	Host     string
	DBName   string
	Port     int
	SSLMode  string
	Password string
	User     string
}

func PostgresURL() string {
	return fmt.Sprintf("host=%s user=%s password=%v dbname=%s port=%d sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
	)
}

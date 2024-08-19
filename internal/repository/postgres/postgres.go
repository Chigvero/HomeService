package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBName"`
	SSLMode  string `yaml:"SSLMode"`
}

const (
	userTable  = "users"
	houseTable = "houses"
	flatTable  = "flats"
)

func NewConnPostgres(cfg Config) (*sqlx.DB, error) {
	connstr := fmt.Sprintf("host=%s port=%s user=%s  dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	conn, err := sqlx.Connect("postgres", connstr)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}

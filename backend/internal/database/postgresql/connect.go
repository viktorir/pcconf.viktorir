package postgresql

import (
	"database/sql"
	"fmt"
	"os"
	"pcconf.viktorir/internal/database"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func Connect() (err error) {
	cfg := Config{
		host:     os.Getenv("PGSQL_HOST"),
		port:     os.Getenv("PGSQL_PORT"),
		user:     os.Getenv("PGSQL_USER"),
		password: os.Getenv("PGSQL_PASSWORD"),
		dbname:   os.Getenv("PGSQL_DBNAME"),
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.host,
		cfg.port,
		cfg.user,
		cfg.password,
		cfg.dbname,
	)
	database.Sql, err = sql.Open("pgx", connString)
	if err != nil {
		return err
	}
	err = database.Sql.Ping()
	if err != nil {
		return err
	}
	return nil
}

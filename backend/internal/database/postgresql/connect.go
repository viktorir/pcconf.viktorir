package postgresql

import (
	"database/sql"
	"fmt"
	"pcconf.viktorir/internal/config"
	"pcconf.viktorir/internal/database"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg config.Database) (err error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
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

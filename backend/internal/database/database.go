package database

import (
	"database/sql"
	"os"
)

var (
	Sql *sql.DB
)

func ReadSQLFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

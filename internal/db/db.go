package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/common/errors"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		errors.LogFatalError(err)
	}

  return db, nil
}

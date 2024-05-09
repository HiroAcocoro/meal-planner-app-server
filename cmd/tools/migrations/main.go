package main

import (
	"os"

	"github.com/HiroAcocoro/meal-planner-app-server/config"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/common/errors"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/db"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPass,
		Addr:                 config.Env.DBAddr,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		errors.LogFatalError(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		errors.LogFatalError(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		errors.LogFatalError(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			errors.LogFatalError(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			errors.LogFatalError(err)
		}
	}
}

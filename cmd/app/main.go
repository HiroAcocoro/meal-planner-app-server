package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/HiroAcocoro/meal-planner-app-server/cmd/api"
	"github.com/HiroAcocoro/meal-planner-app-server/config"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/common/errors"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/db"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPass,
		Addr:                 config.Env.DBAddr,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		errors.LogFatalError(err)
	}

	initMysql(db)

  server := api.NewAPIServer(fmt.Sprintf(":%s", config.Env.APIPort), db)
	if err := server.Run(); err != nil {
		errors.LogFatalError(err)
	}
}

func initMysql(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		errors.LogFatalError(err)
	}

	log.Println("🐬 MySQL successfully connected!")
}

package main

import (
	"database/sql"
	"github.com/Shvoruk/go-api/cmd/api"
	"github.com/Shvoruk/go-api/config"
	"github.com/Shvoruk/go-api/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	mySQL, err := db.NewMySQL(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initDB(mySQL)

	server := api.NewAPIServer(config.Envs.APP_PORT, mySQL)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected")
}

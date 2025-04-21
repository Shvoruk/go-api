// @title           Example Gin + Swagger API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
package main

import (
	"database/sql"
	"log"

	"github.com/Shvoruk/go-api/cmd/api"
	"github.com/Shvoruk/go-api/config"
	"github.com/Shvoruk/go-api/db"
	_ "github.com/Shvoruk/go-api/docs"
	"github.com/go-sql-driver/mysql"
)

func main() {

	// Initialise MySQL connection
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

	// Create and run server
	server := api.NewAPIServer(mySQL)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database: Successfully connected")
}

package api

import (
	"database/sql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"

	"github.com/Shvoruk/go-api/auth"
	"github.com/Shvoruk/go-api/service/animal"
	"github.com/Shvoruk/go-api/service/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db *sql.DB
}

func NewAPIServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Run() error {

	router := gin.Default()

	// Configure CORS middleware.
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // update with your client origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	animalRepo := animal.NewRepo(s.db)
	userRepo := user.NewRepo(s.db)

	animalHandler := animal.NewHandler(animalRepo)
	authHandler := auth.NewHandler(userRepo)

	animalHandler.RegisterRoutes(router)
	authHandler.RegisterRoutes(router)

	// serve swagger UI at /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080"+"/swagger/doc.json"), // <-- point to the JSON
	))
	return router.Run()
}

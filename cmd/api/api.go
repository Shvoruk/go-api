package api

import (
	"database/sql"
	"github.com/Shvoruk/go-api/service/animal"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Run() error {
	router := gin.Default()
	animalRepo := animal.NewRepo(s.db)
	animalHandler := animal.NewHandler(animalRepo)
	animalHandler.RegisterRoutes(router)
	return router.Run(":" + s.addr)
}

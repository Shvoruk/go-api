package animal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Shvoruk/go-api/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	repo types.AnimalRepo
}

func NewHandler(repo types.AnimalRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/animals", h.handleGetAll)
		api.POST("/animals", h.handleCreate)
		api.GET("/animals/:id", h.handleGetByID)
		api.PUT("/animals/:id", h.handleUpdateByID)
		api.DELETE("/animals/:id", h.handleDeleteByID)
	}
}

func (h *Handler) handleGetAll(c *gin.Context) {
	animals, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Someone messed up"})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, animals)
}

func (h *Handler) handleGetByID(c *gin.Context) {
	id := c.Param("id")
	animal, err := h.repo.Get(id)
	if err != nil || animal == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Animal with ID:%s not found", id),
		})
		return
	}
	c.JSON(http.StatusOK, animal)
}

func (h *Handler) handleCreate(c *gin.Context) {
	var a types.Animal
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	newA, err := h.repo.Create(&a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Someone messed up"})
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, newA)
}

func (h *Handler) handleUpdateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var a types.Animal
	a.ID = id
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	newA, err := h.repo.Update(&a)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Animal with ID:%d not found", id),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Someone messed up"})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, newA)
}

func (h *Handler) handleDeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Animal with ID:%s not found", id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("Animal with ID:%s deleted", id),
	})
}

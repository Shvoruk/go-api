package animal

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Shvoruk/go-api/auth"
	"github.com/Shvoruk/go-api/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo types.AnimalRepo
}

func NewHandler(r types.AnimalRepo) *Handler {
	return &Handler{repo: r}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	public := router.Group("/api/v1")
	{
		public.GET("/animals/:id", h.handleGet)
		public.GET("/animals", h.handleGetAllByCategory)
	}

	protected := router.Group("/api/v1")
	protected.Use(auth.Middleware())
	{
		protected.GET("/animals/my", h.handleGetAllByUser)
		protected.POST("/animals", h.handleCreate)
		protected.DELETE("/animals/:id", h.handleDelete)
	}
}

// handleGet godoc
//
//	@Summary		Retrieve a specific animal by ID
//	@Description	Gets detailed information about an animal using its numeric ID.
//	@Tags			animals
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"Animal ID"
//	@Success		200	{object}	types.Animal		"Successful operation"
//	@Failure		400	{object}	map[string]string	"Invalid id parameter"
//	@Failure		404	{object}	map[string]string	"Animal not found"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/animals/{id} [get]
func (h *Handler) handleGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	a, err := h.repo.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusNotFound)
			return
		}
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if a == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, a)
}

// handleGetAllByCategory godoc
//
//	@Summary		Retrieve animals by category
//	@Description	Retrieves a list of animals filtered by the specified category.
//	@Tags			animals
//	@Accept			json
//	@Produce		json
//	@Param			category	query		string				false	"Animal category"
//	@Success		200			{array}		types.Animal		"List of animals"
//	@Failure		500			{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/animals [get]
func (h *Handler) handleGetAllByCategory(c *gin.Context) {
	category := c.Query("category")
	animals, err := h.repo.GetAllByCategory(category)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, animals)
}

// handleGetAllByUser godoc
//
//	@Summary		Retrieve authenticated user's animals
//	@Description	Retrieves a list of animals that belong to the authenticated user.
//	@Tags			animals
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{array}		types.Animal		"List of user's animals"
//	@Failure		401	{object}	map[string]string	"Unauthorized"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/animals/my [get]
func (h *Handler) handleGetAllByUser(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}
	animals, err := h.repo.GetAllByUser(userID.(int))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, animals)
}

// handleCreate godoc
//
//	@Summary		Create a new animal record
//	@Description	Creates a new animal record for the authenticated user.
//	@Tags			animals
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			animal	body		types.Animal		true	"Animal data"
//	@Success		201		{object}	map[string]int		"ID of the newly created animal"
//	@Failure		400		{object}	map[string]string	"Invalid request body"
//	@Failure		401		{object}	map[string]string	"Unauthorized"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/animals [post]
func (h *Handler) handleCreate(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}
	var a types.Animal
	if err := c.ShouldBindJSON(&a); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	// Assign the user ID from context
	a.UserID = userID.(int)

	newID, err := h.repo.Create(&a)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": newID})
}

// handleDelete godoc
//
//	@Summary		Delete an animal record
//	@Description	Deletes an animal record if it belongs to the authenticated user.
//	@Tags			animals
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	int	true	"Animal ID"
//	@Success		204	"No Content"
//	@Failure		400	{object}	map[string]string	"Invalid id parameter"
//	@Failure		401	{object}	map[string]string	"Unauthorized"
//	@Failure		404	{object}	map[string]string	"Animal not found"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/animals/{id} [delete]
func (h *Handler) handleDelete(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}
	animalID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err = h.repo.Delete(userID.(int), animalID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusNotFound)
			return
		}
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

package auth

import (
	"github.com/Shvoruk/go-api/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Handler struct {
	repo types.UserRepo
}

func NewHandler(r types.UserRepo) *Handler {
	return &Handler{repo: r}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/auth/signup", h.handleSignUp)
		api.POST("/auth/login", h.handleLogIn)
	}
}

// handleSignUp godoc
//
//	@Summary		Register a new user
//	@Description	Registers a new user, hashes the password, creates a JWT token, and returns it.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		types.SignUpRequest	true	"User signup data"
//	@Success		201		{object}	map[string]string	"Token generated"
//	@Failure		400		{object}	map[string]string	"Bad request or email already exists"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/auth/signup [post]
func (h *Handler) handleSignUp(c *gin.Context) {
	var req types.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	e, err := h.repo.ExistsByEmail(req.Email)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if e {
		c.Status(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// Create user model
	var user types.User
	user.Username = req.Username
	user.Email = req.Email
	user.Password = string(hashedPassword)

	// Insert into DB
	id, err := h.repo.Create(&user)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	user.ID = id

	// Create JWT
	t, err := createToken(user)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": t})
}

// handleLogIn godoc
//
//	@Summary		Authenticate a user
//	@Description	Authenticates a user using email and password, and returns a JWT token.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		types.LoginRequest	true	"User login credentials"
//	@Success		200			{object}	map[string]string	"Token generated"
//	@Failure		400			{object}	map[string]string	"Invalid request body"
//	@Failure		401			{object}	map[string]string	"Unauthorized"
//	@Failure		500			{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/auth/login [post]
func (h *Handler) handleLogIn(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	u, err := h.repo.GetByEmail(req.Email)
	if err != nil || u == nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	token, err := createToken(*u)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

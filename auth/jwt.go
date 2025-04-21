package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"

	"github.com/Shvoruk/go-api/config"
	"github.com/Shvoruk/go-api/types"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(config.Envs.JWTSecret)

func createToken(u types.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  u.ID,
			"username": u.Email,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		},
	)
	return token.SignedString(secret)
}

func verifyToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

// Middleware checks for a valid JWT in the "Authorization" header
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		// Assumes format "Bearer <token>"
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		// Verify the token
		token, err := verifyToken(tokenString)
		if err != nil || !token.Valid {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		// We expect token.Claims to be jwt.MapClaims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		userIDVal, ok := claims["user_id"]
		if !ok {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		// Typically, user_id is numeric in the token (=> float64 in MapClaims)
		floatID, ok := userIDVal.(float64)
		if !ok {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		userID := int(floatID)
		// Store userID in context
		c.Set("user_id", userID)
		// Next handler
		c.Next()
	}
}

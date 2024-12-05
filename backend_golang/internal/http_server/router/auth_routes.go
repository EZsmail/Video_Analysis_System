package router

import (
	"net/http"

	"backend-golang/internal/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")

	authGroup.GET("/google/login", func(c *gin.Context) {
		url := auth.AuthURL()
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	authGroup.GET("/google/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code not provided"})
			return
		}

		userInfo, err := auth.ExchangeCode(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userID := userInfo["id"].(string)
		token, err := auth.GenerateJWT(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "user": userInfo})
	})
}

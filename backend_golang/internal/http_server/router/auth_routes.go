package router

import (
	"fmt"
	"net/http"

	"backend-golang/internal/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterAuthRoutes(router *gin.Engine, log *zap.Logger) {
	authGroup := router.Group("/login")

	// Google OAuth Login
	authGroup.GET("/", func(c *gin.Context) {
		// url generation for google auth
		url := auth.AuthURL()
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	// google callback
	authGroup.GET("/redirect", func(c *gin.Context) {
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

		// get id from google api
		userID := userInfo["id"].(string)

		// create session
		err = auth.CreateSession(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
			return
		}

		// TODO: Add logging
		// c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": userInfo})

		log.Info("Redirect succesfull", zap.String("user info", fmt.Sprintf("%+v", userInfo)))

		c.Redirect(http.StatusPermanentRedirect, "http://localhost:8080/")
	})

	// Logout (удаление сессии)
	authGroup.POST("/logout", func(c *gin.Context) {
		err := auth.DestroySession(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to destroy session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	})
}

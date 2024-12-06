package router

import (
	"backend-golang/internal/http_server/middleware"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Broker interface {
	SendTask(string, []byte) error
}

type StatusUpdater interface {
	InsertStatus(string, string) error
	GetStatus(string) (string, error)
}

type ResultUpdater interface {
	InsertResult(context.Context, string, string) error
	GetResult(context.Context, string) (string, error)
}

// TODO: Change to interface
func SetupRouter(logger *zap.Logger, broker Broker, mongo ResultUpdater, pg StatusUpdater, debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.GinLogger(logger))

	// TODO: Change logger

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "all good"})
	})

	private := r.Group("/private")
	private.Use(middleware.SessionMiddleware())

	RegisterAuthRoutes(r, logger)
	RegisterUploadRoutes(private, logger, broker, pg)
	RegisterStatusRoutes(private, logger, pg)
	RegisterResultRoutes(r, logger, pg, mongo)

	return r
}

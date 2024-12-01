package router

import (
	"backend-golang/internal/middleware"
	"backend-golang/internal/router/routes"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func SetupRouter(logger *zap.Logger, conn *amqp.Connection, debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.GinLogger(logger))

	routes.RegisterUploadRoutes(r, logger, conn)
	routes.RegisterResultRoutes(r, logger)
	routes.RegisterStatusRoutes(r, logger)

	return r
}

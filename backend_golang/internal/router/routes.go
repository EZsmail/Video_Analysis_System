package router

import (
	"backend-golang/internal/middleware"
	"backend-golang/internal/storage/mongo"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Sender interface {
	SendTask(string, []byte) error
}

type StatusUpdater interface {
	InsertStatus(context.Context, string, string) error
	GetStatus(context.Context, string) (string, error)
}

type ResultUpdater interface {
}

func SetupRouter(logger *zap.Logger, broker Sender, db *mongo.MongoDB, debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.GinLogger(logger))

	RegisterUploadRoutes(r, logger, broker, db)
	RegisterStatusRoutes(r, logger, db)
	RegisterResultRoutes(r, logger)

	return r
}

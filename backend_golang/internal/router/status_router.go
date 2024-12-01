package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StatusGetter interface {
	GetStatus(context.Context, string) (string, error)
}

func RegisterStatusRoutes(r *gin.Engine, logger *zap.Logger, db StatusGetter) {
	r.GET("/status/:processing_id", func(c *gin.Context) {
		processingID := c.Param("processing_id")
		status, err := db.GetStatus(c.Request.Context(), processingID)
		if err != nil {
			logger.Error("Failed to get status", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"processing_id": processingID, "status": status})
		logger.Info("Status retrieved", zap.String("processing_id", processingID))
	})
}

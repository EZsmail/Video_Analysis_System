package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TODO: Add context
type StatusGetter interface {
	GetStatus(string) (string, error)
}

func RegisterStatusRoutes(r *gin.RouterGroup, logger *zap.Logger, pg StatusGetter) {
	r.GET("/status/:processing_id", func(c *gin.Context) {
		processingID := c.Param("processing_id")

		status, err := pg.GetStatus(processingID)
		if err != nil {
			logger.Error("Failed to get status", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"processing_id": processingID, "status": status})
		logger.Info("Status retrieved", zap.String("processing_id", processingID))
	})
}

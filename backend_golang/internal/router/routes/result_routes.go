package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterResultRoutes(r *gin.Engine, logger *zap.Logger) {
	r.GET("/result/:processing_id", func(c *gin.Context) {
		processingID := c.Param("processing_id")
		filePath := "./results/" + processingID + ".csv"

		c.File(filePath)
		logger.Info("Result served", zap.String("processing_id", processingID))
	})
}

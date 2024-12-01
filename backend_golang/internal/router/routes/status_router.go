package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterStatusRoutes(r *gin.Engine, logger *zap.Logger) {
	r.GET("/status/:processing_id", func(c *gin.Context) {
		processingID := c.Param("processing_id")
		// Заглушка: возвращаем статус "в обработке"
		c.JSON(http.StatusOK, gin.H{"processing_id": processingID, "status": "in_progress"})
		logger.Info("Status checked", zap.String("processing_id", processingID))
	})
}

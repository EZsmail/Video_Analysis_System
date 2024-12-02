package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResultGetter interface {
	GetResult(context.Context, string) (string, error)
}

func RegisterResultRoutes(r *gin.Engine, logger *zap.Logger, statusDB StatusGetter, resultDB ResultGetter) {
	r.GET("/result/:processing_id", func(c *gin.Context) {
		processingID := c.Param("processing_id")

		// Проверяем статус обработки
		status, err := statusDB.GetStatus(processingID)
		if err != nil {
			logger.Error("Failed to get status", zap.Error(err), zap.String("processing_id", processingID))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve status"})
			return
		}

		if status != "completed" {
			logger.Info("Result not ready", zap.String("processing_id", processingID), zap.String("status", status))
			c.JSON(http.StatusAccepted, gin.H{"message": "Result is not ready", "processing_id": processingID, "status": status})
			return
		}

		// Извлекаем результат из MongoDB
		result, err := resultDB.GetResult(c.Request.Context(), processingID)
		if err != nil {
			logger.Error("Failed to retrieve result", zap.Error(err), zap.String("processing_id", processingID))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve result"})
			return
		}

		if result == "" {
			logger.Warn("Result is empty", zap.String("processing_id", processingID))
			c.JSON(http.StatusNotFound, gin.H{"error": "Result not found"})
			return
		}

		// Отправка результата как CSV-файл
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename="+processingID+".csv")
		c.String(http.StatusOK, result)

		logger.Info("Result served", zap.String("processing_id", processingID))
	})
}

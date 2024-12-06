package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TODO: Change context
type Task struct {
	ProcessingID string `json:"processing_id"`
	FilePath     string `json:"file_path"`
}

type ResultSaver interface {
	InsertStatus(string, string) error
}

func RegisterUploadRoutes(r *gin.RouterGroup, logger *zap.Logger, broker Broker, db ResultSaver) {
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			logger.Error("Failed to get file", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
			return
		}

		processingID := time.Now().Format("20060102150405")
		filePath := "./uploads/" + file.Filename

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			logger.Error("Failed to save file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Save status to PostgreSQL
		if err := db.InsertStatus(processingID, "in_progress"); err != nil {
			logger.Error("Failed to save status", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save status"})
			return
		}

		task := map[string]string{"processing_id": processingID, "file_path": filePath}
		body, _ := json.Marshal(task)
		if err := broker.SendTask("video_processing", body); err != nil {
			logger.Error("Failed to send task to broker", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process task"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"processing_id": processingID})
	})
}

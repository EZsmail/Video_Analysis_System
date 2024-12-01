package routes

import (
	"backend-golang/internal/mq"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Task struct {
	ProcessingID string `json:"processing_id"`
	FilePath     string `json:"file_path"`
}

func RegisterUploadRoutes(r *gin.Engine, logger *zap.Logger, conn *amqp.Connection) {
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			logger.Error("Failed to get file from request", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
			return
		}

		processingID := time.Now().Format("20060102150405")
		filePath := "./uploads/" + file.Filename

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			logger.Error("Failed to save file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		task := Task{ProcessingID: processingID, FilePath: filePath}
		body, _ := json.Marshal(task)
		if err := mq.SendTask(conn, "video_processing", body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process task"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"processing_id": processingID})
	})
}

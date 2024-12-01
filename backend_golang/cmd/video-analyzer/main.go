package main

import (
	"backend-golang/internal/config"
	"backend-golang/internal/logger"
	"backend-golang/internal/mq"
	"backend-golang/internal/router"
	"backend-golang/internal/storage/mongo"
	"backend-golang/internal/storage/postgresql"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// const (
// 	envDev  = "dev"
// 	envProd = "prod"
// )

func main() {
	path := "config/dev.yaml"
	cfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	log, err := logger.InitLogger(cfg.LogPath, cfg.Debug)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Debug("logger started")
	defer log.Sync()

	rabbitConn, err := mq.ConnectRabbitMQ(cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatal("connection rabbitmq failed", zap.Error(err))
	}
	defer rabbitConn.Close()

	postgresDB, err := postgresql.ConnectPostgreSQL(
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.User,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Database,
		cfg.PostgreSQL.TableStatus,
		log,
	)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL", zap.Error(err))
	}
	defer postgresDB.DB.Close()

	db, err := mongo.ConnectMongoDB(cfg.MongoDB.URL, cfg.MongoDB.Database, cfg.MongoDB.CollectionStatus, cfg.MongoDB.CollectionResult, log)
	if err != nil {
		log.Fatal("connection mongodb failed", zap.Error(err))
	}

	r := router.SetupRouter(log, rabbitConn, db, cfg.Debug)
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Info("Starting server on :8080")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("start server failed", zap.Error(err))
	}
}

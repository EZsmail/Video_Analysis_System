package main

import (
	"backend-golang/internal/config"
	"backend-golang/internal/http_server/router"
	"backend-golang/internal/logger"
	"backend-golang/internal/mq"
	mongo "backend-golang/internal/storage/mongodb"
	"backend-golang/internal/storage/postgresql"
	"log"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var debug = logger.IsDebug()

	log, err := logger.InitLogger(cfg.LogPath, debug)
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

	pg, err := postgresql.ConnectPostgreSQL(
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
	defer pg.Close()

	mongo, err := mongo.ConnectMongoDB(cfg.MongoDB.URL, cfg.MongoDB.Database, cfg.MongoDB.CollectionResult, cfg.MongoDB.CollectionResult, log)
	if err != nil {
		log.Fatal("connection mongodb failed", zap.Error(err))
	}

	r := router.SetupRouter(log, rabbitConn, mongo, pg, debug)

	log.Info("Starting server on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("start server failed", zap.Error(err))
	}
}

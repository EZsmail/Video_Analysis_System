package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoDB struct {
	Client           *mongo.Client
	Database         string
	CollectionResult *mongo.Collection
}

// mongodb connect
func ConnectMongoDB(url, database, collectionStatus, collectionResults string, logger *zap.Logger) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}

	logger.Info("Connected to MongoDB", zap.String("url", url))
	return &MongoDB{
		Client:           client,
		Database:         database,
		CollectionResult: client.Database(database).Collection(collectionResults),
	}, nil
}

// insert result by id
func (db *MongoDB) InsertResult(ctx context.Context, processingID, resultPath string) error {
	_, err := db.CollectionResult.InsertOne(ctx, map[string]string{
		"processing_id": processingID,
		"result_path":   resultPath,
	})
	return err
}

// get result by id
func (db *MongoDB) GetResult(ctx context.Context, processingID string) (string, error) {
	var result struct {
		ResultPath string `bson:"result_path"`
	}
	err := db.CollectionResult.FindOne(ctx, map[string]string{"processing_id": processingID}).Decode(&result)
	return result.ResultPath, err
}

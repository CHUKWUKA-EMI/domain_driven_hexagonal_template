package db

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to the database
func ConnectDB(dbURL string, logger *logrus.Entry) (*mongo.Client, error) {
	return connectMongoDB(dbURL, logger)
}

func connectMongoDB(dbURL string, logger *logrus.Entry) (*mongo.Client, error) {
	logger.Info("Connecting to mongodb")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbURL).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, err
	}

	logger.Info("Successfully connected to MongoDB!")

	return client, nil
}

package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/scratchingmycranium/golang-rest/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type utils struct {
	config *config.Config
}

func InitUtils(cfg *config.Config) *utils {
	return &utils{
		config: cfg,
	}
}

func (u *utils) InitMongo() *mongo.Client {
	// Set up the MongoDB client options
	clientOptions := options.Client().ApplyURI(u.config.MongoUri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

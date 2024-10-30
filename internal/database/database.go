// database/database.go
package database

import (
	"Timeline/internal/helpers"
	"Timeline/internal/logger"
	"Timeline/internal/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client   *mongo.Client
	database *mongo.Database
	Messages *helpers.CollectionHelper[*models.MessageEvent]
}

func Initialize(uri string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(5 * time.Second).
		SetConnectTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		// Attempt to disconnect if ping fails
		disconnectCtx, disconnectCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer disconnectCancel()
		_ = client.Disconnect(disconnectCtx) // Best effort disconnect

		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database("timeline")
	l := logger.GetLogger()
	l.Debug("Connected to MongoDB database: %s", db.Name())

	messageConfig := helpers.CollectionConfig{
		Name: "message_create",
		Indexes: []mongo.IndexModel{
			{Keys: bson.D{{Key: "message_id", Value: 1}}},
			{Keys: bson.D{{Key: "guild_id", Value: 1}}},
			{Keys: bson.D{{Key: "channel_id", Value: 1}}},
			{Keys: bson.D{{Key: "created_at", Value: -1}}},
		},
	}

	messages, err := helpers.NewCollectionHelper[*models.MessageEvent](db, messageConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create message collection: %w", err)
	}

	return &Database{
		client:   client,
		database: db,
		Messages: messages,
	}, nil
}

func (db *Database) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

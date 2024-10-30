package helpers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type CollectionConfig struct {
	Name    string
	Indexes []mongo.IndexModel
}

type CollectionHelper[T any] struct {
	collection *mongo.Collection
}

func NewCollectionHelper[T any](db *mongo.Database, config CollectionConfig) (*CollectionHelper[T], error) {
	coll := db.Collection(config.Name)

	if len(config.Indexes) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		opts := options.CreateIndexes().SetMaxTime(2 * time.Minute)
		_, err := coll.Indexes().CreateMany(ctx, config.Indexes, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to create indexes: %w", err)
		}
	}

	return &CollectionHelper[T]{
		collection: coll,
	}, nil
}

func (h *CollectionHelper[T]) InsertOne(ctx context.Context, document T) error {
	// Handle BaseModel fields if the document implements the interface
	if model, ok := any(document).(interface{ GetBase() *BaseModel }); ok {
		base := model.GetBase()
		now := time.Now()
		if base.CreatedAt.IsZero() {
			base.CreatedAt = now
		}
		base.UpdatedAt = now
	}

	opts := options.InsertOne().SetComment(fmt.Sprintf("Insert %T", document))
	_, err := h.collection.InsertOne(ctx, document, opts)
	return err
}

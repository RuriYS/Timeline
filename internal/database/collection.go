// collections.go
package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection[T any] struct {
	collection *mongo.Collection
}

func (c *Collection[T]) CreateIndexes(ctx context.Context, indexes []mongo.IndexModel) error {
	opts := options.CreateIndexes().SetMaxTime(2 * time.Minute)
	_, err := c.collection.Indexes().CreateMany(ctx, indexes, opts)
	return err
}

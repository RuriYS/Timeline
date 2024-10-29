package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	GuildID   string             `bson:"guild_id"`
	ChannelID string             `bson:"channel_id"`
	UserID    string             `bson:"user_id,omitempty"`
	EventType string             `bson:"event_type"`
	Content   string             `bson:"content,omitempty"`
	Metadata  interface{}        `bson:"metadata,omitempty"`
	Timestamp time.Time          `bson:"timestamp"`
}

type Database struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func Initialize(uri string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("timeline")
	collection := db.Collection("events")

	// Create indexes
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "guild_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "channel_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "user_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "event_type", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "timestamp", Value: -1}},
		},
		// Compound index for common queries
		{
			Keys: bson.D{
				{Key: "guild_id", Value: 1},
				{Key: "channel_id", Value: 1},
				{Key: "timestamp", Value: -1},
			},
		},
	}

	// Create indexes with modern options
	opts := options.CreateIndexes().SetMaxTime(2 * time.Minute)
	_, err = collection.Indexes().CreateMany(ctx, indexes, opts)
	if err != nil {
		return nil, err
	}

	return &Database{
		client:     client,
		database:   db,
		collection: collection,
	}, nil
}

func (d *Database) StoreEvent(ctx context.Context, event *Event) error {
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	// Use ordered option for better error handling
	opts := options.InsertOne().SetComment("Store timeline event")
	_, err := d.collection.InsertOne(ctx, event, opts)
	return err
}

func (d *Database) GetEvents(ctx context.Context, filter map[string]interface{}) ([]*Event, error) {
	findOptions := options.Find().
		SetSort(bson.D{{Key: "timestamp", Value: -1}}).
		SetComment("Get timeline events")

	if limit, ok := filter["limit"].(int64); ok {
		findOptions.SetLimit(limit)
		delete(filter, "limit")
	}

	query := bson.M{}
	for key, value := range filter {
		switch key {
		case "since":
			if t, ok := value.(time.Time); ok {
				query["timestamp"] = bson.M{"$gte": t}
			}
		case "until":
			if t, ok := value.(time.Time); ok {
				if query["timestamp"] == nil {
					query["timestamp"] = bson.M{}
				}
				query["timestamp"].(bson.M)["$lte"] = t
			}
		default:
			query[key] = value
		}
	}

	cursor, err := d.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []*Event
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (d *Database) DeleteEvents(ctx context.Context, filter map[string]interface{}) error {
	query := bson.M{}
	for key, value := range filter {
		query[key] = value
	}

	opts := options.Delete().SetComment("Delete timeline events")
	_, err := d.collection.DeleteMany(ctx, query, opts)
	return err
}

func (d *Database) UpdateEvent(ctx context.Context, id primitive.ObjectID, update map[string]interface{}) error {
	updateDoc := bson.M{"$set": update}
	opts := options.Update().SetComment("Update timeline event")

	_, err := d.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		updateDoc,
		opts,
	)
	return err
}

// GetEventStats returns statistics about events for a guild
func (d *Database) GetEventStats(ctx context.Context, guildID string, since time.Time) (map[string]int64, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"guild_id":  guildID,
			"timestamp": bson.M{"$gte": since},
		}}},
		{{Key: "$group", Value: bson.M{
			"_id":   "$event_type",
			"count": bson.M{"$sum": 1},
		}}},
	}

	opts := options.Aggregate().SetComment("Get event statistics")
	cursor, err := d.collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []struct {
		EventType string `bson:"_id"`
		Count     int64  `bson:"count"`
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	stats := make(map[string]int64)
	for _, result := range results {
		stats[result.EventType] = result.Count
	}
	return stats, nil
}

func (d *Database) Close(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}

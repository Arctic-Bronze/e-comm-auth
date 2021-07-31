package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	database = "local-e-commerce-app"
)

var (
	uri = os.Getenv("MONGO_URI")
)

var client *mongo.Client

func init() {
	ctx, _ := GetCtx()
	opt := options.Client().ApplyURI(uri).SetMaxPoolSize(1).SetDirect(true)
	client, _ = mongo.Connect(ctx, opt)
}

func GetCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

func GetCollection(collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_USER     = "abdullahgungor"
	DB_PASSWORD = "Ag7410"
)

func GetDB() *mongo.Database {
	databaseURL := "mongodb+srv://" + DB_USER + ":" + DB_PASSWORD + "@cluster0.xbwcqpz.mongodb.net/?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		panic(err)
	}

	db := client.Database("productdb")

	return db

}

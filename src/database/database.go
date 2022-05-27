package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Export struct{}

var contextTimeout = 10 * time.Second

func (c Export) ConnectMongo() (*mongo.Client, error) {
	urlDb := "mongodb://root:rootpassword@206.189.38.110:27017/?authSource=admin&readPreference=primary&ssl=false"

	client, err := mongo.NewClient(options.Client().ApplyURI(urlDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), contextTimeout)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

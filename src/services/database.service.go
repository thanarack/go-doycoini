package services

// import (
// 	"context"
// 	"log"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// func (c Export) ListDatabaseNames() ([]string, error) {
// 	databases, _ := connection.ConnectMongo()

// 	ctx, err := context.WithTimeout(context.Background(), contextTimeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer databases.Disconnect(ctx)

// 	list, _ := databases.ListDatabaseNames(ctx, bson.M{})

// 	return list, nil
// }

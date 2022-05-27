package services

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"koicoini.com/sdk/v2/src/models"
)

func (c Export) ListTokens() ([]models.Tokens, error) {
	// Connect to databases
	databases, _ := connection.ConnectMongo()
	defer databases.Disconnect(ctx)

	// Looking data on tokens collection.
	coll := databases.Database(databaseName).Collection(collectionToken)
	filter := bson.D{{}}

	// Find all tokens. and returned []bytes
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		log.Panic(err)
	}

	// Definitely data is returned
	var data []models.Tokens

	// Convert data to bson M and append to data.
	for cursor.Next(ctx) {
		// Decode individual result
		var result bson.Raw
		cursor.Decode(&result)
		// Destruction and append the result to the data.
		tokenAttribute := models.Tokens{
			ID:           result.Lookup("_id").ObjectID(),
			Token:        result.Lookup("token").StringValue(),
			Symbol:       result.Lookup("symbol").StringValue(),
			Change24:     result.Lookup("change24").Decimal128(),
			PriceBinance: result.Lookup("price_binance").Decimal128(),
			PriceBitkub:  result.Lookup("price_bitkub").Decimal128(),
			PriceUpbit:   result.Lookup("price_upbit").Decimal128(),
			RoomId:       result.Lookup("room_id").StringValue(),
			CreatedAt:    result.Lookup("created_at").Time(),
			UpdatedAt:    result.Lookup("updated_at").Time(),
		}
		data = append(data, tokenAttribute)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return data, nil
}

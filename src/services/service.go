package services

import (
	"context"

	"koicoini.com/sdk/v2/src/database"
)

type Export struct{}

var connection database.Export
var databaseName = "koicoini"
var collectionToken = "tokens"
var ctx = context.TODO()

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tokens struct {
	ID           primitive.ObjectID   `json:"id"`
	Token        string               `json:"token"`
	Symbol       string               `json:"symbol"`
	Change24     primitive.Decimal128 `json:"change24"`
	PriceBinance primitive.Decimal128 `json:"price_binance"`
	PriceBitkub  primitive.Decimal128 `json:"price_bitkub"`
	PriceUpbit   primitive.Decimal128 `json:"price_upbit"`
	RoomId       string               `json:"room_id"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
}

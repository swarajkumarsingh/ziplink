package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlModel struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	LongUrl string             `json:"longUrl" bson:"longUrl"`
	ShortId string             `json:"shortId" bson:"shortId"`
	Expiry  time.Time          `json:"expiry" bson:"expiry"`
}

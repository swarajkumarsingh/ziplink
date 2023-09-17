package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type UrlModel struct {
	Id bson.ObjectId 	`json:"id" bson:"_id"`
	LongUrl string		`json:"longUrl" bson:"longUrl"`
	ShortId string		`json:"shortId" bson:"shortId"`
	Expiry  time.Time	`json:"expiry" bson:"expiry"`
}
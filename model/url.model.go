package model

import (
	"time"
)

type UrlModel struct {
	LongUrl string        `json:"longUrl" bson:"longUrl"`
	ShortId string        `json:"shortId" bson:"shortId"`
	Expiry  time.Duration `json:"expiry" bson:"expiry"`
}

package controller

import "time"

type Response struct {
	Url         string
	CustomShort string
	Expiry      time.Duration
}

type Request struct {
	LongUrl string `json:"url"`
}
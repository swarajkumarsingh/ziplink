package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/ziplink/functions/general"
	redis "github.com/swarajkumarsingh/ziplink/infra/redis"
)

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error":   true,
		"message": message,
	})
}

func GetShortId(longUrl string) (string, error) {
	if longUrl == "" || general.IsNotValidURL(longUrl) {
		return "", errors.New("invalid long url")
	}

	counter, err := redis.IncrementCounter()
	if err != nil {
		return "", nil
	}

	shortId := general.ConvertToBase64ID(counter)

	if shortId == "" {
		return "", errors.New("internal server error")
	}

	return shortId, nil
}

func CacheLongUrl(shortId string, longUrl string) error {

	if general.IsNotValidURL(longUrl) {
		fmt.Println("invalid url, for caching")
	}

	var ttl = time.Hour * 24
	err := redis.Set(shortId, longUrl, ttl)

	if err != nil {
		return err
	}

	return nil
}

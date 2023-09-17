package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/ziplink/conf"
	"github.com/swarajkumarsingh/ziplink/functions/general"
	redisUtils "github.com/swarajkumarsingh/ziplink/infra/redis"
	"gopkg.in/mgo.v2"
)

type UrlController struct {
	session *mgo.Session
}

type Response struct {
	Url string
	CustomShort string
	Expiry time.Duration
}

func NewUrlController(s *mgo.Session) *UrlController {
	return &UrlController{s}
}

func (uc UrlController) CreateUrl(c *gin.Context) {

	// create redis client

	// create mongodb client 

	// get a counter from redis
	counter, err := redisUtils.IncrementCounter()
	if err != nil {
		panic(err)
	}

	// convert the counter number to b64 encoder chr length 7
	shortId := general.ConvertToBase64ID(counter)

	var response Response = Response{
		Url: longUrl,
		CustomShort: shortId,
		Expiry: conf.FreedomRedisTTL,
	}

	// Save in mongodb

	// Save in Redis cache with TTL(1 Day)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "po=",
	})
}

func (uc UrlController) RedirectUrl(c *gin.Context) {
	// create mongodb client 

	// find url from shortUrlId

	// redirect to longURL

	c.Redirect(http.StatusPermanentRedirect, "https://google.com/")
}
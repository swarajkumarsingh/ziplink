package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type UrlController struct {
	session *mgo.Session
}

func NewUrlController(s *mgo.Session) *UrlController {
	return &UrlController{s}
}

func (uc UrlController) CreateUrl(c *gin.Context) {

	// create redis client

	// create mongodb client 

	// get a counter from redis

	// convert the counter number to b64 encoder chr length 7

	// Set expiry

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
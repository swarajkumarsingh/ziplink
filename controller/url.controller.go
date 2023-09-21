package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/ziplink/conf"
	"github.com/swarajkumarsingh/ziplink/functions/general"
	"github.com/swarajkumarsingh/ziplink/functions/logger"
	"github.com/swarajkumarsingh/ziplink/infra/db"
	redis "github.com/swarajkumarsingh/ziplink/infra/redis"
	"github.com/swarajkumarsingh/ziplink/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUrl(c *gin.Context) {
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil || body.LongUrl == "" || !general.IsValidURL(body.LongUrl) {
		SendErrorResponse(c, http.StatusBadRequest, "Invalid url found")
		return
	}

	shortId, err := GetShortId()
	if err != nil {
		fmt.Println("1" +err.Error())
		SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	var content = model.UrlModel{
		LongUrl: body.LongUrl,
		ShortId: shortId,
		Expiry:  conf.FreedomRedisTTL,
	}

	msg, err := db.InsertUrl(content)
	if err != nil {
		fmt.Println("2" +err.Error())
		SendErrorResponse(c, http.StatusInternalServerError, msg+err.Error())
		return
	}

	err = CacheLongUrl(shortId, content.LongUrl)
	if err != nil {
		logger.WithRequest(c).Errorln(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  content,
	})
}

func RedirectUrl(c *gin.Context) {
	shortId := c.Param("url")
	if shortId == "" || len(shortId) != 7 {
		SendErrorResponse(c, http.StatusBadRequest, "Invalid shortUrl")
		return
	}

	val, err := redis.Get(shortId)
	if err != nil || val == "" || !general.IsValidURL(val) {
		fmt.Println("url not found in cache")

		// fetch from DB
		urlModel, err := db.FindOne(shortId)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				SendErrorResponse(c, http.StatusNotFound, "Specific Url not found")
				return
			}
			SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		if !general.IsValidURL(urlModel.LongUrl) {
			SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.Redirect(http.StatusPermanentRedirect, urlModel.LongUrl)
		return
	}

	fmt.Println("Redirecting from cache")
	c.Redirect(http.StatusPermanentRedirect, val)
}

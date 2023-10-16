package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swarajkumarsingh/ziplink/controller"
	"github.com/swarajkumarsingh/ziplink/infra/db"
	redis "github.com/swarajkumarsingh/ziplink/infra/redis"
)

var version string = "1.0"

func enableCORS() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    c.Writer.Header().Set("Access-Control-Allow-Headers",
      "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Api-Key, token, User-Agent, Referer")
    c.Writer.Header().Set("AllowCredentials", "true")
    c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

    if c.Request.Method == "OPTIONS" {
      return
    }

    c.Next()
  }
}

func main() {
  gin.SetMode(gin.ReleaseMode)
  
  r := gin.Default()
  r.Use(enableCORS())
  r.Use(gin.Recovery())

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  db.Init()
  redis.Init()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "health ok",
    })
  })

  r.POST("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "health ok",
    })
  })

  r.GET("/:url", controller.RedirectUrl)
  r.POST("/create-url", controller.CreateUrl)

  log.Printf("Server Started, version: %s", version)
  r.Run(":8080")
}

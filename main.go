package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/ziplink/controller"
	"github.com/swarajkumarsingh/ziplink/db"
	"github.com/swarajkumarsingh/ziplink/router"
)

var version string = "1.0"

func startLog() gin.HandlerFunc {
  return func(c *gin.Context) {
    u := ""
    reqLog := map[string]string{
      "requestID": u,
      "startTime": time.Now().Format("2006-01-02 15:04:05.000000000"),
      "endTime":   "",
    }
    c.Set("reqLog", reqLog)
    c.Next()
  }
}

func EnableCORS() gin.HandlerFunc {
  return func(c *gin.Context) {
    // corsEnabledStr := vault.GetOrDefaultString(c, vault.KeyCORSEnabled, "false")
    // corsAllowedDomains := strings.ToLower(vault.GetOrDefaultString(c, vault.KeyCORSOriginDomains, ""))

    // if origin := strings.ToLower(c.GetHeader("Origin")); corsEnabledStr == "true" && origin != "" && (strings.Contains(corsAllowedDomains, "*") || strings.Contains(corsAllowedDomains, origin)) {

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
  r := gin.Default()

  r.Use(startLog())
  r.Use(EnableCORS())

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "health ok",
    })
  })

  controller.NewUrlController(db.GetSession())
  router.SetupUrlRoutes(r)

  log.Printf("Server Started, version: %s", version)
  r.Run("localhost:8080")
}

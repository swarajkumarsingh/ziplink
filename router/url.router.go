package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/ziplink/controller"
	"github.com/swarajkumarsingh/ziplink/db"
)

func SetupUrlRoutes(r *gin.Engine) {
  uc := controller.NewUrlController(db.GetSession())
  userRouter := r.Group("/")
  {
    userRouter.POST("/create-url", uc.CreateUrl)
    userRouter.GET("/:url", uc.RedirectUrl)
  }
}

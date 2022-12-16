package router

import (
	"github.com/gin-gonic/gin"
	"go-rss-reader-service/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	rss := r.Group("/rss")
	{
		rss.POST("/reader", controllers.ReadRssFeed)
	}
	return r
}

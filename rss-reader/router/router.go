package router

import (
	"github.com/gin-gonic/gin"
	"go-rss-reader-service/controllers"
	"go-rss-reader-service/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	rss := r.Group("/rss")
	{
		rss.POST("/reader", controllers.ReadRssFeed)
		rss.POST("/url/add", controllers.AddRssUrl)
		rss.GET("/url/list", controllers.GetAllRssUrls)
	}
	return r
}

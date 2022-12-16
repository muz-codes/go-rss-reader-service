package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rss-reader-service/dto"
	services "go-rss-reader-service/service"
	"go-rss-reader-service/utils"
	"net/http"
)

func ReadRssFeed(ctx *gin.Context) {
	var rssReaderRequest dto.RssReaderRequest
	var rssReaderResponse dto.RssFeedResponse
	if err := ctx.ShouldBindJSON(&rssReaderRequest); err != nil {
		if fieldErrorsArray := utils.GetFieldErrorsArray(err); fieldErrorsArray != nil {
			ctx.JSON(http.StatusBadRequest, fieldErrorsArray)
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rssItems, err := services.RssReader(rssReaderRequest)
	if err != nil {
		rssReaderResponse.Success = false
		rssReaderResponse.Message = fmt.Sprintf("error in rssFeed")
		ctx.JSON(http.StatusInternalServerError, rssReaderResponse)
		return
	}
	rssReaderResponse.Success = true
	rssReaderResponse.Message = fmt.Sprintf("rss feed successful")
	rssReaderResponse.RssItems = rssItems
	ctx.JSON(http.StatusOK, rssReaderResponse)
}

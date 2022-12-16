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
	return
}

func AddRssUrl(ctx *gin.Context) {
	var addRssUrlRequest dto.AddRssUrlRequest
	var addRssUrlResponse dto.AddRssUrlResponse
	if err := ctx.ShouldBindJSON(&addRssUrlRequest); err != nil {
		if fieldErrorsArray := utils.GetFieldErrorsArray(err); fieldErrorsArray != nil {
			ctx.JSON(http.StatusBadRequest, fieldErrorsArray)
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	addedRssUrl, err := services.AddRssUrl(&addRssUrlRequest)
	if err != nil {
		addRssUrlResponse.Success = false
		addRssUrlResponse.Message = fmt.Sprintf("add rss url failed : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, addRssUrlResponse)
		return
	}
	addRssUrlResponse.Success = true
	addRssUrlResponse.Message = fmt.Sprintf("rss url successfully added")
	addRssUrlResponse.Url = addedRssUrl
	ctx.JSON(http.StatusOK, addRssUrlResponse)
	return
}

func GetAllRssUrls(ctx *gin.Context) {
	var getRssUrlRequest dto.GetRssUrlRequest
	var getAllRssUrlsResponse dto.GetAllRssUrlsResponse
	if err := ctx.ShouldBindQuery(&getRssUrlRequest); err != nil {
		if fieldErrorsArray := utils.GetFieldErrorsArray(err); fieldErrorsArray != nil {
			ctx.JSON(http.StatusBadRequest, fieldErrorsArray)
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rssUrlList, totalCount, totalPages, err := services.GetRssUrls(&getRssUrlRequest)
	if err != nil {
		getAllRssUrlsResponse.Success = false
		getAllRssUrlsResponse.Message = fmt.Sprintf("failed to get rss urls: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, getAllRssUrlsResponse)
		return
	}
	getAllRssUrlsResponse.Success = true
	getAllRssUrlsResponse.Message = fmt.Sprintf("%v rss urls fetchd", len(rssUrlList))
	getAllRssUrlsResponse.Urls = rssUrlList
	getAllRssUrlsResponse.TotalCount = totalCount
	getAllRssUrlsResponse.TotalPages = totalPages
	ctx.JSON(http.StatusOK, getAllRssUrlsResponse)
	return
}

func UpdateUrl(ctx *gin.Context) {
	var updateRssUrlRequest dto.UpdateRssUrlRequest
	var updateRssUrlResponse dto.UpdateRssUrlResponse
	if err := ctx.ShouldBindJSON(&updateRssUrlRequest); err != nil {
		if fieldErrorsArray := utils.GetFieldErrorsArray(err); fieldErrorsArray != nil {
			ctx.JSON(http.StatusBadRequest, fieldErrorsArray)
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedRssUrl, err := services.UpdateRssUrl(&updateRssUrlRequest)
	if err != nil {
		updateRssUrlResponse.Success = false
		updateRssUrlResponse.Message = fmt.Sprintf("update rss url failed : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, updateRssUrlResponse)
		return
	}
	updateRssUrlResponse.Success = true
	updateRssUrlResponse.Message = fmt.Sprintf("rss url updated successfuly")
	updateRssUrlResponse.Url = updatedRssUrl
	ctx.JSON(http.StatusOK, updateRssUrlResponse)
	return
}

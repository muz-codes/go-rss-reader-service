package service

import (
	"errors"
	"go-rss-reader-service/db_utils"
	"go-rss-reader-service/dto"
	"go-rss-reader-service/utils"
	"go.uber.org/zap"
)

var logger = zap.L()

func AddRssUrl(addRssUrlRequest *dto.AddRssUrlRequest) (dto.RssUrl, error) {
	var rssUrlDto dto.RssUrl
	// check if url already exists in db
	exist, err := utils.CheckIfUrlExistInDb(addRssUrlRequest.Url)
	if err != nil {
		logger.Error("error in AddRssUrl", zap.Error(err))
		return dto.RssUrl{}, err
	}
	if exist {
		errorOccurred := errors.New("url already exist in db")
		logger.Error("error in AddRssUrl", zap.Error(errorOccurred))
		return dto.RssUrl{}, errorOccurred
	}

	// validate url by calling the url
	_, err = utils.ValidateUrl(addRssUrlRequest.Url)
	if err != nil {
		logger.Error("error in AddRssUrl", zap.Error(err))
		return dto.RssUrl{}, err
	}

	rssUrl, err := db_utils.AddRssUrlToDb(addRssUrlRequest.Url)
	if err != nil {
		logger.Error("error in AddRssUrl", zap.Error(err))
		return dto.RssUrl{}, err
	}
	// assigning rss Url model data to DTO for organized response
	rssUrlDto.Id = rssUrl.ID
	rssUrlDto.Url = rssUrl.Url
	return rssUrlDto, nil
}

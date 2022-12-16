package service

import (
	"errors"
	"go-rss-reader-service/db_utils"
	"go-rss-reader-service/dto"
	"go-rss-reader-service/utils"
	"go.uber.org/zap"
)

func GetRssUrls(getRssUrlRequest *dto.GetRssUrlRequest) ([]dto.RssUrl, int64, int64, error) {
	if err := GenerateParamsFromGetRssUrlRequest(getRssUrlRequest); err != nil {
		return nil, 0, 0, err
	}
	rssUrlDto, err := db_utils.GetRssUrls(getRssUrlRequest)
	if err != nil {
		logger.Error("error in GetRssUrls", zap.Error(err))
		return nil, 0, 0, err
	}
	totalCount, totalPages, err := GetUrlCountAndPageCount(getRssUrlRequest)
	if err != nil {
		logger.Error("error in GetRssUrls", zap.Error(err))
		return nil, 0, 0, err
	}

	return rssUrlDto, totalCount, totalPages, nil
}

func GetUrlCountAndPageCount(getRssUrlRequest *dto.GetRssUrlRequest) (int64, int64, error) {
	totalCount, err := db_utils.GetRssUrlsCount()
	if err != nil {
		logger.Error("error in GetRssUrls", zap.Error(err))
		return 0, 0, err
	}
	totalPages := utils.CalculateTotalPagesForPagination(totalCount, getRssUrlRequest.Limit)
	return totalCount, totalPages, err
}

func GenerateParamsFromGetRssUrlRequest(getRssUrlRequest *dto.GetRssUrlRequest) error {
	if getRssUrlRequest.Limit < 1 {
		getRssUrlRequest.Limit = 20
	}
	if getRssUrlRequest.Page < 1 {
		getRssUrlRequest.Page = 1
	}
	if getRssUrlRequest.Sort == "" {
		getRssUrlRequest.Sort = "created_at desc"
	} else if getRssUrlRequest.Sort != "created_at desc" && getRssUrlRequest.Sort != "created_at asc" {
		errorOccurred := errors.New("invalid sorting")
		logger.Error("error in GenerateParamsFromGetRssUrlRequest", zap.Error(errorOccurred))
		return errorOccurred
	}
	return nil
}

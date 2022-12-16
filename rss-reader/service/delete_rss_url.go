package service

import (
	"go-rss-reader-service/db_utils"
	"go-rss-reader-service/dto"
	"go.uber.org/zap"
)

func DeleteRssUrl(deleteRssUrlRequest *dto.DeleteRssUrlRequest) (dto.RssUrl, error) {
	rssUrlDeleted, err := db_utils.DeleteRssUrlById(deleteRssUrlRequest.Id)
	if err != nil {
		logger.Error("error in DeleteRssUrl", zap.Error(err))
		return dto.RssUrl{}, err
	}
	return rssUrlDeleted, nil
}

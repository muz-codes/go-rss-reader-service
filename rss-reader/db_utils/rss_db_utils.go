package db_utils

import (
	"errors"
	"fmt"
	"go-rss-reader-service/db"
	"go-rss-reader-service/dto"
	"go-rss-reader-service/models"
	"go.uber.org/zap"
)

var logger = zap.L()

func AddRssUrlToDb(url string) (models.RssUrl, error) {
	var rssUrl models.RssUrl
	rssUrl.Url = url
	if err := db.DbConnection.Create(&rssUrl).Error; err != nil {
		logger.Error("error in AddRssUrlToDb", zap.Error(err))
		return models.RssUrl{}, err
	}
	return rssUrl, nil
}

func GetRssUrls(getRssUrlRequest *dto.GetRssUrlRequest) ([]dto.RssUrl, error) {
	offset := int((getRssUrlRequest.Page - 1) * getRssUrlRequest.Limit)
	var rssUrl []models.RssUrl
	var rssUrlDto []dto.RssUrl
	if err := db.DbConnection.Offset(offset).Limit(int(getRssUrlRequest.Limit)).Order(getRssUrlRequest.Sort).Find(&rssUrl).Scan(&rssUrlDto).Error; err != nil {
		return nil, err
	}
	return rssUrlDto, nil
}

func GetRssUrlsCount() (int64, error) {
	var rssUrl models.RssUrl
	var rssUrlsCount int64
	if err := db.DbConnection.Find(&rssUrl).Count(&rssUrlsCount).Error; err != nil {
		return 0, err
	}
	return rssUrlsCount, nil
}

func GetRssUrlByUrl(url string) (dto.RssUrl, error) {
	var rssUrl models.RssUrl
	var rssUrlDto dto.RssUrl
	if err := db.DbConnection.Where("url = ?", url).First(&rssUrl).Scan(&rssUrlDto).Error; err != nil {
		logger.Error("error in GetRssUrlByUrl", zap.Error(err))
		return dto.RssUrl{}, err
	}
	return rssUrlDto, nil
}

func UpdateRssUrl(id int64, url string) (models.RssUrl, error) {
	var rssUrl models.RssUrl
	rssUrl.Url = url
	if err := db.DbConnection.Where("id = ?", id).Updates(rssUrl).Error; err != nil {
		logger.Error("error in UpdateRssUrl", zap.Error(err))
		return models.RssUrl{}, err
	}
	return rssUrl, nil
}

func DeleteRssUrlById(id int64) (dto.RssUrl, error) {
	var rssUrlDto dto.RssUrl
	query := db.DbConnection.Table("rss_urls").Where("id = ?", id).Scan(&rssUrlDto).Delete(&models.RssUrl{})
	if query.Error != nil {
		logger.Error("error in DeleteRssUrlById", zap.Error(query.Error))
		return rssUrlDto, query.Error
	}
	if query.RowsAffected == 0 {
		errorOccurred := errors.New(fmt.Sprintf("no url found to delete with id %v", id))
		return rssUrlDto, errorOccurred
	}
	return rssUrlDto, nil
}

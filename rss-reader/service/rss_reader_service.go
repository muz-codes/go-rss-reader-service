package service

import (
	"go.uber.org/zap"
	// importing go-rss-reader package
	reader "github.com/muzamilQP/go-rss-reader"
	"go-rss-reader-service/dto"
)

func RssReader(rssReaderRequest dto.RssReaderRequest) ([]reader.RssItem, error) {
	logger := zap.L()
	var rssItem reader.RssItem
	allFeedItems, err := rssItem.Parse(rssReaderRequest.RssUrls)
	if err != nil {
		logger.Error("error in RssReader Service", zap.Error(err))
		return nil, err
	}
	return allFeedItems, nil
}

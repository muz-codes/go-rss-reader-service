package dto

import (
	reader "github.com/muzamilQP/go-rss-reader"
)

type RssFeedResponse struct {
	Success  bool             `json:"success"`
	Message  string           `json:"message"`
	RssItems []reader.RssItem `json:"items"`
}

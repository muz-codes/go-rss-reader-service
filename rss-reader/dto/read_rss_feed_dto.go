package dto

type RssReaderRequest struct {
	RssUrls []string `json:"rss_urls" binding:"required"`
}

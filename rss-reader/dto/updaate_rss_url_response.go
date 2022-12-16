package dto

type UpdateRssUrlResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Url     RssUrl `json:"rss_url"`
}

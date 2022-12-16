package dto

type AddRssUrlResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Url     RssUrl `json:"rss_url"`
}

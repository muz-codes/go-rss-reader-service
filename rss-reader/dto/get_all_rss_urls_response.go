package dto

type GetAllRssUrlsResponse struct {
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Urls       []RssUrl `json:"rss_url_list"`
	TotalCount int64    `json:"total_count"`
	TotalPages int64    `json:"total_pages"`
}

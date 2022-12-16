package dto

type RssUrl struct {
	Id  uint   `json:"id"`
	Url string `json:"url"`
}

type AddRssUrlRequest struct {
	Url string `json:"url" binding:"required"`
}

type GetRssUrlRequest struct {
	Page  int64  `json:"page" binding:"required"`
	Limit int64  `json:"limit" binding:"required"`
	Sort  string `json:"sort"`
}

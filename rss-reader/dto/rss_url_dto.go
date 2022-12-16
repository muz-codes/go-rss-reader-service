package dto

type RssUrl struct {
	Id  uint   `json:"id"`
	Url string `json:"url"`
}

type AddRssUrlRequest struct {
	Url string `json:"url" binding:"required"`
}

type GetRssUrlRequest struct {
	Page  int64  `form:"page" binding:"required"`
	Limit int64  `form:"limit" binding:"required"`
	Sort  string `form:"sort"`
}

type UpdateRssUrlRequest struct {
	Id  int64  `json:"id" binding:"required"`
	Url string `json:"url" binding:"required"`
}

type DeleteRssUrlRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

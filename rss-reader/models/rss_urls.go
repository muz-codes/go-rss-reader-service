package models

import "gorm.io/gorm"

type RssUrl struct {
	gorm.Model
	Url string `json:"url" gorm:"column:url"`
}

func (r *RssUrl) TableName() string {
	return "rss_urls"
}

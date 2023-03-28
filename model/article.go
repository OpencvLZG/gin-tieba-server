package model

import "time"

type Article struct {
	ArticleId      int64     `json:"id" xorm:"pk autoincr"`
	UserId         int64     `json:"user_id"`
	UserName       string    `json:"user_name" xorm:"varchar(50) not null"`
	Belong         string    `json:"belong" xorm:"varchar(20) notnull"`
	ArticleLike    int64     `json:"like"`
	Title          string    `json:"title" binding:"required" xorm:"varchar(50) notnull"`
	ArticleContext string    `json:"article_context" binding:"required" xorm:"varchar(500) notnull"`
	ImgUrl         string    `json:"img_url" xorm:"varchar(150)"`
	CreateTime     time.Time `json:"create_time" xorm:"created"`
	UpdateTime     time.Time `json:"update_time" xorm:"updated"`
}

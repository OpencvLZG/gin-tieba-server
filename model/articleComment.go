package model

import "time"

type ArticleComment struct {
	Id         int64     `json:"id" xorm:"autoincr pk notnull"`
	ArticleId  int64     `json:"article_id"  binding:"required" xorm:"not null"`
	UserId     int64     `json:"user_id" xorm:"not null"`
	CommentId  int64     `json:"comment_id"  xorm:"autoincr"`
	Context    string    `json:"context" binding:"required" xorm:"varchar(500) notnull"`
	IsDel      uint8     `json:"is_del" xorm:"default 0"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	DelTime    time.Time `json:"del_time" xorm:"updated"`
}

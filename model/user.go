package model

import "time"

type User struct {
	Id         int64     `json:"id" xorm:"autoincr pk"`
	Avatar     string    `json:"avatar" xorm:"varchar(100)"`
	Account    string    `json:"account" xorm:"varchar(25)"`
	Username   string    `json:"username"binding:"required" xorm:"varchar(50) notnull unique"`
	Secret     string    `json:"secret" xorm:"varchar(125) notnull"`
	Password   string    `json:"password" binding:"required"xorm:"varchar(125) notnull"`
	IsBlocked  string    `json:"is_blocked" xorm:"varchar(10) default('false')"`
	Email      string    `json:"email" xorm:"varchar(50)"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}

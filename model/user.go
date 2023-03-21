package model

import "time"

type User struct {
	Id         int64     `json:"id" xorm:"autoincr pk"`
	Username   string    `json:"username" xorm:"varchar(50) notnull unique"`
	Secret     string    `json:"secret" xorm:"varchar(125) notnull"`
	Password   string    `json:"password" xorm:"varchar(125) notnull"`
	Email      string    `json:"email" xorm:"varchar(50)"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}

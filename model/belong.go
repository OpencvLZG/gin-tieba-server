package model

import "time"

type Belong struct {
	BelongId   int64     `json:"belong_id" xorm:"pk autoincr"`
	UserId     int64     `json:"user_id" xorm:"notnull"`
	IsDel      uint8     `json:"is_del" xorm:"default 0"`
	BelongName string    `json:"belong_name" xorm:"varchar(25) notnull unique"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}

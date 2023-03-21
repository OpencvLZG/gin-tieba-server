package model

import "time"

type Belong struct {
	BelongId        int64  `json:"belong_id" xorm:"pk autoincr"`
	UserId          int64  `json:"user_id" xorm:"notnull"`
	BelongName      string `json:"belong_name" xorm:"varchar(25) notnull"`
	BelongFollowers int64
	CreateTime      time.Time `json:"create_time" xorm:"created"`
	UpdateTime      time.Time `json:"update_time" xorm:"update"`
}

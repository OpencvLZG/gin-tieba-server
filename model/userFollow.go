/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/25/2023
  @desc: //TODO
**/

package model

import "time"

type UserFollow struct {
	Id          int64     `json:"id" xorm:"autoincr pk not null"`
	UserId      int64     `json:"user_id" xorm:"not null unique(un_index)"`
	FollowId    int64     `json:"follow_id" xorm:"not null unique(un_index)"`
	CreatedTime time.Time `json:"created_time" xorm:"created"`
}

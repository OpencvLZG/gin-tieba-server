/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/31/2023
  @desc: //TODO
**/

package dao

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	BelongFollowDao struct {
	}
)

func (b *BelongFollowDao) BelongFollow(belongFollow *model.BelongFollower) error {
	orm := util.Orm
	_, err := orm.InsertOne(belongFollow)
	return err
}

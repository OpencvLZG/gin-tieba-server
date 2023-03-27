/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/25/2023
  @desc: //TODO
**/

package dao

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	BelongDao struct {
	}
)

func (b *BelongDao) CreateBelong(belong *model.Belong) error {
	orm := util.Orm
	_, err := orm.InsertOne(belong)
	return err
}

func (b *BelongDao) DelBelong(belong *model.Belong) error {
	orm := util.Orm
	belong.IsDel = 1
	_, err := orm.Cols("is_del").
		Update(belong, &model.Belong{UserId: belong.UserId, BelongId: belong.BelongId})
	return err
}

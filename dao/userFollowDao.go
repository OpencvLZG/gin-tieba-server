/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/29/2023
  @desc: //TODO
**/

package dao

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	UserFollowDao struct {
	}
)

func (uf *UserFollowDao) FollowUser(userFollow *model.UserFollow) error {
	orm := util.Orm
	_, err := orm.InsertOne(userFollow)
	return err
}

func (uf *UserFollowDao) GetFollowedUid(userFollow *model.UserFollow) error {
	orm := util.Orm
	_, err := orm.Get(userFollow)
	return err

}

func (uf *UserFollowDao) UnFollowUser(userFollow *model.UserFollow) error {
	orm := util.Orm
	_, err := orm.Delete(userFollow)
	return err
}

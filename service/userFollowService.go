/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/31/2023
  @desc: //TODO
**/

package service

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
)

type (
	UserFollowService struct {
	}
)

func (uf *UserFollowService) FollowUser(userFollow *model.UserFollow) error {
	userFollowDao := new(dao.UserFollowDao)
	err := userFollowDao.FollowUser(userFollow)
	return err
}

func (uf *UserFollowService) UnFollowUser(userFollow *model.UserFollow) error {
	userFollowDao := new(dao.UserFollowDao)
	err := userFollowDao.UnFollowUser(userFollow)
	return err
}

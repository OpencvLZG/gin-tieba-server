/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/16
  @desc: //TODO
**/

package service

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
)

type (
	UserManageService struct {
	}
)

func (a *UserManageService) BlockerUser(userId int64) error {
	user := new(model.User)
	user.Id = userId
	user.IsBlocked = "ture"
	userManageDao := new(dao.UserManageDao)
	err := userManageDao.BlockedUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *UserManageService) UnBlockerUser(userId int64) error {
	user := new(model.User)
	user.Id = userId
	user.IsBlocked = "false"
	userManageDao := new(dao.UserManageDao)
	err := userManageDao.BlockedUser(user)
	if err != nil {
		return err
	}
	return nil
}

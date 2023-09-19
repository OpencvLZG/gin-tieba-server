/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/16
  @desc: //TODO
**/

package dao

import (
	"errors"
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	UserManageDao struct {
	}
)

func (u *UserManageDao) QueryUsers(offset int) (*[]model.User, error) {
	orm := util.Orm
	userList := make([]model.User, 0)
	err := orm.Limit(10, offset).OrderBy("create_time").Find(&userList)
	if err != nil {
		return &userList, errors.New("查询失败,数据库查询失败")
	}
	return &userList, err
}

func (a *UserManageDao) BlockedUser(user *model.User) error {
	orm := util.Orm
	affect, err := orm.Where("id = ?", user.Id).Update(user)
	if err != nil {
		return err
	} else if affect == 0 {
		return errors.New("冻结失败")
	}
	return nil
}

func (u *UserManageDao) UnBlockedUser(user *model.User) error {
	orm := util.Orm
	affect, err := orm.Where("id = ?", user.Id).Update(user)
	if err != nil {
		return err
	} else if affect == 0 {
		return errors.New("解冻失败")
	}
	return nil
}

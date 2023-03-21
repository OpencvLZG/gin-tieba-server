package dao

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	UserDao struct {
	}
)

func (u *UserDao) Insert(user *model.User) error {
	orm := util.Orm
	_, err := orm.InsertOne(user)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserDao) Login(user *model.User) (bool, error) {
	var res bool
	orm := util.Orm
	res, err := orm.Get(user)
	return res, err
}

func (u *UserDao) QueryBySecret(user *model.User) error {
	orm := util.Orm
	_, err := orm.Where("username = ?", user.Username).Get(user)
	return err
}

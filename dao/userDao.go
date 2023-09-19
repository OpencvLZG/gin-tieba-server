package dao

import (
	"fmt"
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
	return err
}

func (u *UserDao) LoginByUsername(user *model.User) (bool, error) {
	var res bool
	orm := util.Orm
	res, err := orm.Get(user)
	return res, err
}
func (u *UserDao) LoginByAccount(user *model.User) (bool, error) {
	var res bool
	orm := util.Orm
	res, err := orm.Get(user)
	return res, err
}
func (u *UserDao) GetUserByUid(user *model.User) error {
	orm := util.Orm
	_, err := orm.Get(user)
	return err
}

func (u *UserDao) QueryBySecret(user *model.User) (bool, error) {
	orm := util.Orm
	res, err := orm.Where("username = ?", user.Username).Get(user)
	return res, err
}

func (u *UserDao) ReName(user *model.User) error {
	fmt.Printf("%v", user)
	orm := util.Orm
	_, err := orm.Where("id = ?", user.Id).Update(user)
	if err != nil {
		return err
	}
	return nil
}

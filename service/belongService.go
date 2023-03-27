/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/25/2023
  @desc: //TODO
**/

package service

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
)

type (
	BelongService struct {
	}
)

func (b *BelongService) CreatBelong(belong *model.Belong) error {
	belongDao := new(dao.BelongDao)
	err := belongDao.CreateBelong(belong)
	return err
}

func (b *BelongService) DelBelong(belong *model.Belong) error {
	belongDao := new(dao.BelongDao)
	err := belongDao.DelBelong(belong)
	return err
}

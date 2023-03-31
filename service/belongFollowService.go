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
	BelongFollowService struct {
	}
)

func (b *BelongFollowService) BelongFollow(belongFollow *model.BelongFollower) error {

	belongFollowDao := new(dao.BelongFollowDao)
	err := belongFollowDao.BelongFollow(belongFollow)
	if err != nil {
		return err
	}
	return err
}

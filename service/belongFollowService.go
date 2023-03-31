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
	"errors"
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
)

type (
	BelongFollowService struct {
	}
)

func (b *BelongFollowService) BelongFollow(belongFollow *model.BelongFollower) error {
	belongDao := new(dao.BelongDao)
	belong := new(model.Belong)
	belong.BelongId = belongFollow.BelongId
	res, err := belongDao.SearchBelongById(belong)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("不存在该贴吧，关注失败")
	}
	belongFollowDao := new(dao.BelongFollowDao)
	err = belongFollowDao.BelongFollow(belongFollow)
	if err != nil {
		return err
	}
	return err
}

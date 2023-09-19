/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/31/2023
  @desc: //TODO
**/

package controller

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
)
import . "ginFlutterBolg/util"

type (
	BelongFollowController struct {
	}
)

func (b *BelongFollowController) BelongFollow(c *gin.Context) {
	belongFollow := new(model.BelongFollower)
	if err := c.Bind(&belongFollow); err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	belongDao := new(dao.BelongDao)
	belong := new(model.Belong)
	belong.BelongId = belongFollow.BelongId
	res, err := belongDao.SearchBelongById(belong)
	if err != nil {
		return
	}
	if res != true {
		ResponseStatusOk(c, 400, "贴吧不存在", nil)
		return
	}
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseStatusOk(c, 400, "获取用户信息失败", err.Error())
		return
	}

	belongFollow.UserId = user.Id
	//print(user.Id)
	belongFollowService := new(service.BelongFollowService)
	err = belongFollowService.BelongFollow(belongFollow)
	if err != nil {
		ResponseStatusOk(c, 400, "关注失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "关注成功", nil)

}

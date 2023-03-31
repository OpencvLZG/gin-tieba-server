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
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
)

import . "ginFlutterBolg/util"

type (
	UserFollowController struct {
	}
)

func (u *UserFollowController) FollowUser(c *gin.Context) {
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseStatusOk(c, 400, "数据获取错误", err.Error())
		return
	}
	userFollow := new(model.UserFollow)
	if err := c.Bind(&userFollow); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	userFollow.UserId = user.Id
	userFollowService := new(service.UserFollowService)
	err = userFollowService.FollowUser(userFollow)
	if err != nil {
		return
	}
	ResponseStatusOk(c, 200, "关注成功", "")
}

/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/16
  @desc: //TODO
**/

package controller

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
import . "ginFlutterBolg/util"

type (
	UserManageController struct {
	}
)

func (umc *UserManageController) GetUserList(c *gin.Context) {
	offset := c.Query("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	userManageDao := new(dao.UserManageDao)
	userList, err := userManageDao.QueryUsers(offsetInt)
	if err != nil {
		ResponseStatusOk(c, 400, "数据查询失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "获取成功", gin.H{
		"list": userList,
	})

}

func (a *UserManageController) BlockedUser(c *gin.Context) {
	userId := c.Query("user_id")
	intUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ResponseStatusFail(c, 400, "表单绑定失败", err)
		return
	}
	userManageService := new(service.UserManageService)
	err = userManageService.BlockerUser(intUserId)
	if err != nil {
		ResponseStatusFail(c, 400, "冻结失败", err)
		return
	}
	ResponseStatusOk(c, 200, "冻结成功", err)
}

func (a *UserManageController) UnBlockedUser(c *gin.Context) {
	userId := c.Query("user_id")
	intUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ResponseStatusFail(c, 400, "表单绑定失败", err)
		return
	}
	userManageService := new(service.UserManageService)
	err = userManageService.BlockerUser(intUserId)
	if err != nil {
		ResponseStatusFail(c, 400, "解冻失败", err)
		return
	}
	ResponseStatusOk(c, 200, "解冻成功", err)
}

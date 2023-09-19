/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/25/2023
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
	BelongController struct {
	}
)

func (b *BelongController) CreateBelong(c *gin.Context) {
	belong := new(model.Belong)
	if err := c.Bind(&belong); err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseStatusOk(c, 400, "用户数据获取失败", err.Error())
		return
	}
	belong.UserId = user.Id
	belongService := new(service.BelongService)
	err = belongService.CreatBelong(belong)
	if err != nil {
		ResponseStatusOk(c, 400, "创建失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "创建成功", "")
}

func (b *BelongController) DelBelong(c *gin.Context) {
	belong := new(model.Belong)
	if err := c.Bind(&belong); err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	belongService := new(service.BelongService)
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseStatusOk(c, 400, "用户数据获取失败", err.Error())
		return
	}
	belong.UserId = user.Id
	err = belongService.DelBelong(belong)
	if err != nil {
		ResponseStatusOk(c, 400, "文章删除错误", nil)
		return
	}
	ResponseStatusOk(c, 200, "删除成功", nil)
}

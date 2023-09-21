/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/18
  @desc: //TODO
**/

package controller

import (
	. "ginFlutterBolg/util"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	AdminController struct {
	}
)

func (a *AdminController) GenerateTemporaryTokens(c *gin.Context) {
	privilege, exist := c.Get("privilege")
	if !exist {
		ResponseStatusFail(c, 400, "登录信息过期请重新登录", exist)
		return

	} else if privilege != "admin" {
		ResponseStatusFail(c, 400, "权限不足无法创建", false)
		return
	}

	currentTime := time.Now()
	accessToken := new(AccessToken)
	//		获取当前时间
	accessToken.Secret = currentTime.Format("2006-01-02 15-04-05")
	//		令牌七天后失效
	accessToken.Expire = time.Now().Add(7 * 24 * time.Hour).Unix()
	err := accessToken.GenerateToken()
	if err != nil {
		ResponseStatusFail(c, 400, "生成令牌失败", err.Error())
	}
	ResponseStatusOk(c, 200, "生成成功", accessToken.Token)

}

func (a *AdminController) GetLogTotal(c *gin.Context) {

}

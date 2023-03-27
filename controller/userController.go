package controller

import (
	"errors"
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
import . "ginFlutterBolg/util"

type (
	UserController struct {
	}
)

func (h *UserController) UserLogin(c *gin.Context) {
	loginUser := new(service.LoginUser)
	if err := c.Bind(&loginUser); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}

	userService := new(service.UserService)
	res, err := userService.Login(loginUser)
	if err != nil {
		ResponseStatusOk(c, 400, "登录失败", res)
		return
	}
	ResponseStatusOk(c, 200, "登录成功", res)
}
func (h *UserController) UserRegister(c *gin.Context) {
	registerUser := new(service.RegisterUser)
	if err := c.Bind(&registerUser); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	userService := new(service.UserService)
	err := userService.Register(registerUser)
	if err != nil {
		ResponseStatusOk(c, 400, "注册过程出错", err.Error())
		return
	}
	ResponseStatusOk(c, 400, "注册成功", "")

}

func (u *UserController) GetUserInfo(c *gin.Context) (*model.User, error) {
	//获取用户信息
	val, _ := c.Get("userSecret")
	secret := val.(string)
	if secret == "" {
		return nil, errors.New("用户登陆失败")
	}

	userService := new(service.UserService)
	user, err := userService.FindBySecret(secret)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserController) GetUserByUid(c *gin.Context) {
	uid := c.Query("Uid")
	uidInt, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	userService := new(service.UserService)
	user, err := userService.GetUserByUid(uidInt)
	if err != nil {
		ResponseStatusOk(c, 400, "数据查找失败", user)
		return
	}
	userInfo := map[string]interface{}{
		"name":       user.Username,
		"createTime": user.CreateTime,
		"avatar":     user.Avatar,
	}
	ResponseStatusOk(c, 200, "查找成功", userInfo)
}

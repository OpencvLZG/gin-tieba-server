package service

import (
	"errors"
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
	"time"
)

type (
	UserService struct {
	}
	LoginUser struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	LoginResponse struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
	RegisterUser struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"    binding:"required"`
	}
)

func (u *UserService) Login(loginUser *LoginUser) (*LoginResponse, error) {
	userdao := new(dao.UserDao)

	user := new(model.User)
	user.Username = loginUser.Username
	user.Password = util.MD5(loginUser.Password)
	res, err := userdao.Login(user)
	if err != nil {
		return nil, err
	}

	if res != true {
		return nil, errors.New("登录失败")
	}

	accessToken := new(util.AccessToken)
	accessToken.Secret = user.Username
	err = accessToken.GenerateToken()
	if err != nil {
		return nil, err
	}

	loginResponse := new(LoginResponse)
	loginResponse.Username = user.Username
	loginResponse.Token = accessToken.Token
	loginResponse.Email = user.Email

	return loginResponse, nil
}

func (u *UserService) Register(registerUser *RegisterUser) error {
	now := time.Now()

	userdao := new(dao.UserDao)
	user := new(model.User)

	user.Username = registerUser.Username
	user.Password = util.MD5(registerUser.Password)

	user.Email = registerUser.Email
	user.CreateTime = now
	user.UpdateTime = now
	err := userdao.Insert(user)
	if err != nil {
		return err
	}

	return nil

}

func (u *UserService) FindBySecret(secret string) (*model.User, error) {

	//获取用户信息
	userDao := new(dao.UserDao)

	user := new(model.User)
	user.Username = secret
	if err := userDao.QueryBySecret(user); err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, errors.New("用户不存在")
	}

	return user, nil

}

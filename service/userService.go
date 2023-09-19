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
		Account  string `json:"account" binding:"required"`
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
	RegisterResponse struct {
		Id       int64  `json:"id"       binding:"required"`
		Account  string `json:"account"  binding:"required"`
		Username string `json:"username" binding:"required"`
		Email    string `json:"email"    binding:"required"`
	}
)

func (u *UserService) LoginByUsername(loginUser *LoginUser) (*LoginResponse, error) {
	userdao := new(dao.UserDao)

	user := new(model.User)
	user.Username = loginUser.Username
	user.Password = util.MD5(loginUser.Password)
	res, err := userdao.LoginByUsername(user)
	if err != nil {
		return nil, err
	}

	if res != true {
		return nil, errors.New("登录失败")
	}
	if user.IsBlocked == "ture" {
		return nil, errors.New("账号已被管理员锁定，请联系管理员")
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
func (u *UserService) LoginByAccount(loginUser *LoginUser) (*LoginResponse, error) {
	userdao := new(dao.UserDao)

	user := new(model.User)
	user.Username = loginUser.Account
	user.Password = util.MD5(loginUser.Password)
	res, err := userdao.LoginByAccount(user)
	if err != nil {
		return nil, err
	}

	if res != true {
		return nil, errors.New("登录失败")
	}
	if user.IsBlocked == "ture" {
		return nil, errors.New("账号已被管理员锁定，请联系管理员")
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
func (u *UserService) Register(registerUser *RegisterUser) (*RegisterResponse, error) {
	now := time.Now()

	userdao := new(dao.UserDao)
	user := new(model.User)
	user.Account = util.GenerateUID()
	user.Username = registerUser.Username
	user.Password = util.MD5(registerUser.Password)
	user.Secret = util.GenerateKey()
	user.Email = registerUser.Email
	user.CreateTime = now
	user.UpdateTime = now
	err := userdao.Insert(user)
	if err != nil {
		return nil, err
	}
	registerResponse := new(RegisterResponse)
	registerResponse.Id = user.Id
	registerResponse.Account = user.Account
	registerResponse.Username = user.Username
	registerResponse.Email = user.Email
	return registerResponse, nil

}

func (u *UserService) FindBySecret(secret string) (*model.User, error) {

	//获取用户信息
	userDao := new(dao.UserDao)

	user := new(model.User)
	user.Username = secret
	res, err := userDao.QueryBySecret(user)
	if err != nil {
		return nil, err
	}
	if res != true {
		return nil, errors.New("用户不存在")
	}
	return user, nil

}

func (u *UserService) GetUserByUid(uid int64) (*model.User, error) {
	userDao := new(dao.UserDao)
	user := new(model.User)
	user.Id = uid
	err := userDao.GetUserByUid(user)
	return user, err

}

func (u *UserService) ReName(user *model.User) error {
	userDao := new(dao.UserDao)
	err := userDao.ReName(user)
	if err != nil {
		return err
	}
	return nil
}

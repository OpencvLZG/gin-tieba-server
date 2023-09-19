package middle

import "C"
import (
	"bytes"
	"github.com/gin-gonic/gin"
)
import . "ginFlutterBolg/util"

func LoginMiddleWare(context *gin.Context) {
	//	请求头获取Token
	authorization := context.GetHeader("Authorization")
	accessToken := AccessToken{
		Token: authorization,
	}

	secret, err := accessToken.ValidateToken()
	if err != nil {
		ResponseStatusFail(context, 4001, "token验证失败", err.Error())
		context.Abort()
	} else if secret == "" {
		ResponseStatusFail(context, 4002, "secret is empty", "")
		context.Abort()
	}

	//设置用户secret
	context.Set("userSecret", secret)

	context.Next()
	return
}

func AdminValidMiddle(c *gin.Context) {
	// 密钥
	secret := []byte("www.cilang.buzz")
	// 请求头
	authorization := c.GetHeader("Authorization")
	// 判断不等于,则判断是否为临时管理员
	res := bytes.Equal(secret, []byte(authorization))
	if !res {
		accessToken := AccessToken{
			Token: authorization,
		}
		secret, err := accessToken.ValidateToken()
		if err != nil {
			ResponseStatusFail(c, 4001, "管理员密匙校验失败", err.Error())
			c.Abort()
		} else if secret == "" {
			ResponseStatusFail(c, 4002, "secret is empty", "")
			c.Abort()
		}
		c.Set("privilege", "tempAdmin")
	} else {
		c.Set("privilege", "admin")
	}
	c.Next()
}

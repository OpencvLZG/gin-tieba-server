package middle

import "github.com/gin-gonic/gin"
import . "ginFlutterBolg/util"

func LoginMiddleWare(context *gin.Context) {
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

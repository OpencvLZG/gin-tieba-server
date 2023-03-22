package router

import (
	"ginFlutterBolg/controller"
	. "ginFlutterBolg/middle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(gin *gin.Engine) {

	gin.Use(Cors())

	userController := &controller.UserController{}
	articleController := &controller.ArticleController{}
	articleCommentController := &controller.ArticleCommentController{}

	userGroup := gin.Group("/user")
	userGroup.POST("/login", userController.UserLogin)
	userGroup.POST("/register", userController.UserRegister)

	viewGroup := gin.Group("/view")
	viewGroup.GET("list", articleController.GetArticleList)
	viewGroup.GET("search", articleController.SearchArticleById)

	articleGroup := gin.Group("/article")
	articleGroup.Use(LoginMiddleWare)
	articleGroup.POST("create_article", articleController.CreateArticle)

	articleCommentGroup := gin.Group("/articleComment")
	articleCommentGroup.Use(LoginMiddleWare)
	articleCommentGroup.POST("create_comment", articleCommentController.CreateComment)
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

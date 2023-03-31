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
	belongController := &controller.BelongController{}
	belongFollowController := &controller.BelongFollowController{}
	userFollowController := &controller.UserFollowController{}
	userGroup := gin.Group("/user")
	userGroup.POST("/login", userController.UserLogin)
	userGroup.POST("/register", userController.UserRegister)

	viewGroup := gin.Group("/view")
	viewGroup.GET("list", articleController.GetArticleList)
	viewGroup.GET("searchArticle", articleController.SearchArticleById)
	viewGroup.GET("searchArticleByUid", articleController.SearchArticleByUid)
	viewGroup.GET("commentList", articleCommentController.GetArticleCommentList)
	viewGroup.GET("commentListLim", articleCommentController.GetArticleCommentListLimit)
	viewGroup.GET("searchUserByUid", userController.GetUserByUid)

	followGroup := gin.Group("/follow")
	followGroup.Use(LoginMiddleWare)
	followGroup.POST("/followUser", userFollowController.FollowUser)
	followGroup.POST("/unFollowUser", userFollowController.UnFollowUser)
	followGroup.POST("/belongFollow", belongFollowController.BelongFollow)

	articleGroup := gin.Group("/article")
	articleGroup.Use(LoginMiddleWare)
	articleGroup.POST("create_article", articleController.CreateArticle)

	belongGroup := gin.Group("/belong")
	belongGroup.Use(LoginMiddleWare)
	belongGroup.POST("create_belong", belongController.CreateBelong)
	belongGroup.POST("del_belong", belongController.DelBelong)

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

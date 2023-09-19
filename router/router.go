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
	chatController := &controller.ChatController{}
	userManageController := controller.UserManageController{}
	adminController := controller.AdminController{}
	// 用户接口
	userGroup := gin.Group("/user")
	userGroup.POST("/loginByUserName", userController.UserLoginByUserName)
	userGroup.POST("/loginByAccount", userController.UserLoginByAccount)
	userGroup.POST("/register", userController.UserRegister)

	userOpera := userGroup.Group("/opera")
	userOpera.Use(LoginMiddleWare)
	userOpera.POST("/rename", userController.ReName)
	// 用户视图接口
	viewGroup := gin.Group("/view")
	viewGroup.GET("articleList", articleController.GetArticleList)
	viewGroup.GET("articleListOffset", articleController.GetArticleListOffset)
	viewGroup.GET("searchArticle", articleController.SearchArticleById)
	viewGroup.GET("searchArticleByUid", articleController.SearchArticleByUid)
	viewGroup.GET("commentList", articleCommentController.GetArticleCommentList)
	viewGroup.GET("commentListLim", articleCommentController.GetArticleCommentListLimit)
	viewGroup.GET("searchUserByUid", userController.GetUserByUid)
	// 用户关注接口
	followGroup := gin.Group("/follow")
	followGroup.Use(LoginMiddleWare)
	followGroup.POST("/followUser", userFollowController.FollowUser)
	followGroup.POST("/unFollowUser", userFollowController.UnFollowUser)
	followGroup.POST("/belongFollow", belongFollowController.BelongFollow)
	// 用户文章接口
	articleGroup := gin.Group("/article")
	articleGroup.Use(LoginMiddleWare)
	articleGroup.POST("create_article", articleController.CreateArticle)
	// 话题接口
	belongGroup := gin.Group("/belong")
	belongGroup.Use(LoginMiddleWare)
	belongGroup.POST("create_belong", belongController.CreateBelong)
	belongGroup.POST("del_belong", belongController.DelBelong)
	// 文章评论接口
	articleCommentGroup := gin.Group("/articleComment")
	articleCommentGroup.Use(LoginMiddleWare)
	articleCommentGroup.POST("create_comment", articleCommentController.CreateComment)
	// 用户聊天接口
	wsGroup := gin.Group("/ws")
	wsGroup.Use(LoginMiddleWare)
	wsGroup.GET("/chat", chatController.WsHandle)
	// 管理员接口
	manageGroup := gin.Group("/manage")
	manageGroup.Use(AdminValidMiddle)
	manageGroup.POST("/generateToken", adminController.GenerateTemporaryTokens)
	userManage := manageGroup.Group("/user")
	userManage.GET("/getUserList", userManageController.GetUserList)
	userManage.POST("/blockedUser", userManageController.BlockedUser)
	userManage.POST("/unBlockedUser", userManageController.UnBlockedUser)

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

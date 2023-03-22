package controller

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
)
import . "ginFlutterBolg/util"

type (
	ArticleCommentController struct {
	}
)

func (as *ArticleCommentController) CreateComment(c *gin.Context) {
	articleComment := new(model.ArticleComment)
	articleCommentService := new(service.ArticleCommentService)
	userController := new(UserController)
	if err := c.Bind(&articleComment); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseStatusOk(c, 400, "用户登录失败", err.Error())
		return
	}
	articleComment.UserId = user.Id
	err = articleCommentService.CreateComment(articleComment)
	if err != nil {
		ResponseStatusOk(c, 400, "评论失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "评论成功", "")
}

//func (ac * ArticleCommentController) GetArticleCommentList()  {
//	articleCommentService := new(service.ArticleCommentService)
//
//	articleCommentService.GetArticleComment()
//
//}

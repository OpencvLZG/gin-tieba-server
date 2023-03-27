package controller

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
import . "ginFlutterBolg/util"

type (
	ArticleCommentController struct {
	}
)

func (ac *ArticleCommentController) CreateComment(c *gin.Context) {
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

func (ac *ArticleCommentController) GetArticleCommentList(c *gin.Context) {
	articleComment := new(model.ArticleComment)
	articleCommentService := new(service.ArticleCommentService)
	if err := c.Bind(&articleComment); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	articleCommentList, err := articleCommentService.GetArticleComment(articleComment)
	if err != nil {
		ResponseStatusOk(c, 200, "获取失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "获取成功", gin.H{
		"list": articleCommentList,
	})

}

func (ac *ArticleCommentController) GetArticleCommentListLimit(c *gin.Context) {
	offset := c.Query("Offset")
	offsetInt, err := strconv.Atoi(offset)
	articleComment := new(model.ArticleComment)
	articleCommentService := new(service.ArticleCommentService)
	if err := c.Bind(&articleComment); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	articleCommentList, err := articleCommentService.GetArticleCommentLimit(articleComment, offsetInt)
	if err != nil {
		ResponseStatusOk(c, 200, "获取失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "获取成功", gin.H{
		"list": articleCommentList,
	})

}

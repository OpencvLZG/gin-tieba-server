package controller

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
)
import . "ginFlutterBolg/util"

type (
	ArticleController struct {
	}
)

func (a *ArticleController) GetArticleList(c *gin.Context) {
	articleService := new(service.ArticleService)
	list, err := articleService.ArticleList()
	if err != nil {

	}
	ResponseStatusOk(c, 200, "ok", gin.H{
		"list": list,
	})

}

func (a *ArticleController) CreateArticle(c *gin.Context) {
	article := new(model.Article)
	userController := new(UserController)
	articleRequest := new(service.ArticleRequest)
	if err := c.Bind(&articleRequest); err != nil {
		ResponseStatusOk(c, 400, "数据同步错误", err.Error())
		return
	}
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseStatusOk(c, 400, "用户登录失败", err.Error())
		return
	}
	article.UserId = user.Id
	//println(user.Id)
	//println(articleRequest.Title)
	article.Title = articleRequest.Title
	article.ArticleContext = articleRequest.ArticleContext
	article.Belong = articleRequest.Belong
	article.ImgUrl = articleRequest.ImgUrl
	print(article.Title)
	println(articleRequest.Title)
	articleService := new(service.ArticleService)
	err = articleService.InsertArticle(article)
	if err != nil {
		ResponseStatusOk(c, 400, "创建失败，不允许空标题和空内容", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "创建成功", "")
}

func (a *ArticleController) SearchArticleById(c *gin.Context) {
	article := new(model.Article)
	if err := c.Bind(&article); err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	articleService := new(service.ArticleService)
	err := articleService.SearchByArticleId(article)
	print(err)
	if err != nil {
		ResponseStatusOk(c, 410, "搜索失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "搜索成功", article)
}

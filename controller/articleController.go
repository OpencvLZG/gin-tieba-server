package controller

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
		ResponseStatusFail(c, 400, "文章获取失败", err.Error())
	}
	ResponseStatusOk(c, 200, "ok", gin.H{
		"list": list,
	})

}
func (a *ArticleController) GetArticleListOffset(c *gin.Context) {
	page := c.Query("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		ResponseStatusOk(c, 400, "参数查询错误", err.Error())
		return
	}
	articleService := new(service.ArticleService)
	list, err := articleService.GetArticleOffset(intPage)
	if err != nil {
		ResponseStatusFail(c, 400, "文章获取失败", err.Error())
		return
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
	article.UserName = user.Username
	article.Title = articleRequest.Title
	article.ArticleContext = articleRequest.ArticleContext
	article.Belong = articleRequest.Belong
	article.ImgUrl = articleRequest.ImgUrl
	articleService := new(service.ArticleService)
	err = articleService.InsertArticle(article)
	if err != nil {
		ResponseStatusOk(c, 400, "创建失败，不允许空标题和空内容", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "创建成功", nil)
}

func (a *ArticleController) SearchArticleById(c *gin.Context) {
	article := new(model.Article)
	if err := c.Bind(&article); err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	articleService := new(service.ArticleService)
	err := articleService.SearchByArticleId(article)
	//print(err)
	if err != nil {
		ResponseStatusOk(c, 410, "搜索失败", err.Error())
		return
	}
	ResponseStatusOk(c, 200, "搜索成功", article)
}

func (a *ArticleController) SearchArticleByUid(c *gin.Context) {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		ResponseStatusOk(c, 400, "数据绑定错误", err.Error())
		return
	}
	articleService := new(service.ArticleService)
	list, err := articleService.SearchByArticleUid(user.Id)
	if err != nil {

	}
	ResponseStatusOk(c, 200, "ok", gin.H{
		"list": list,
	})

}

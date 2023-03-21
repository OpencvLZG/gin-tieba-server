package service

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
)

type (
	ArticleService struct {
	}
	ArticleRequest struct {
		Belong         string `json:"belong" binding:"required"`
		Title          string `json:"title" binding:"required"`
		ArticleContext string `json:"article_context" binding:"required"`
		ImgUrl         string `json:"img_url"`
	}
	ArticleList struct {
	}
)

func (a *ArticleService) ArticleList() (*[]model.Article, error) {
	articleDao := new(dao.ArticleDao)
	articleList, err := articleDao.GetArticleLimit()
	if err != nil {

	}
	return articleList, err
}

func (a *ArticleService) InsertArticle(article *model.Article) error {
	articleDao := new(dao.ArticleDao)
	err := articleDao.InsertArticle(article)
	if err != nil {

	}
	return nil

}

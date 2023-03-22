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
	if err := articleDao.InsertArticle(article); err != nil {

	}
	return nil
}

func (a *ArticleService) SearchByArticleId(article *model.Article) error {
	articleDao := new(dao.ArticleDao)
	if err := articleDao.GetArticle(article); err != nil {
		return err
	}
	return nil
}

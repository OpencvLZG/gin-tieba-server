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
	return articleList, err
}
func (a *ArticleService) GetArticleOffset(page int) (*[]model.Article, error) {
	articleDao := new(dao.ArticleDao)
	articleList, err := articleDao.GetArticleOffset(page)
	return articleList, err

}
func (a *ArticleService) InsertArticle(article *model.Article) error {
	articleDao := new(dao.ArticleDao)
	err := articleDao.InsertArticle(article)
	return err
}

func (a *ArticleService) SearchByArticleId(article *model.Article) error {
	articleDao := new(dao.ArticleDao)
	err := articleDao.GetArticle(article)
	return err
}

func (a *ArticleService) SearchByArticleUid(uid int64) (*[]model.Article, error) {
	articleDao := new(dao.ArticleDao)
	articleList, err := articleDao.SearchArticleListByUid(uid)
	return articleList, err
}

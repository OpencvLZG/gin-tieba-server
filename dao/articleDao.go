package dao

import (
	"errors"
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	ArticleDao struct {
	}
)

func (a *ArticleDao) GetArticleLimit() (*[]model.Article, error) {
	orm := util.Orm
	articleList := make([]model.Article, 0)
	err := orm.Limit(10).OrderBy("create_time").Find(&articleList)
	if err != nil {

	}
	return &articleList, err

}
func (a *ArticleDao) GetArticleById(article *model.Article) error {
	orm := util.Orm
	res, err := orm.Get(article)
	if err != nil {
		return errors.New("搜索失败,数据库查询失败")
	}
	if res {

	} else {
		return errors.New("搜索失败,没有数据")
	}
	return nil

}
func (a *ArticleDao) InsertArticle(article *model.Article) error {
	orm := util.Orm
	_, err := orm.InsertOne(article)
	return err
}

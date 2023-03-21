package dao

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	ArticleDao struct {
	}
)

func (a *ArticleDao) GetArticleLimit() (*[]model.Article, error) {
	orm := util.Orm
	var err error
	articleList := make([]model.Article, 0)
	err = orm.Limit(10).OrderBy("create_time").Find(&articleList)
	if err != nil {

	}
	return &articleList, err

}
func (a *ArticleDao) InsertArticle(article *model.Article) error {
	orm := util.Orm
	_, err := orm.InsertOne(article)
	return err
}

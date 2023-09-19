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
		return &articleList, errors.New("搜索文章列表错误，数据库搜索错误")
	}
	return &articleList, err

}
func (a *ArticleDao) GetArticleOffset(page int) (*[]model.Article, error) {
	orm := util.Orm
	articleList := make([]model.Article, 0)
	err := orm.Limit(10, page*10).OrderBy("create_time").Find(&articleList)
	if err != nil {
		return &articleList, errors.New("搜索文章列表错误，数据库搜索错误")
	}
	return &articleList, err
}
func (a *ArticleDao) GetArticle(article *model.Article) error {
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
	if err != nil {
		return errors.New("插入文章错误,数据库插入文章错误")
	}
	return err
}

func (a *ArticleDao) SearchArticleListByUid(uid int64) (*[]model.Article, error) {
	orm := util.Orm
	articleList := make([]model.Article, 0)
	err := orm.Where("id = ?", uid).Limit(10).OrderBy("create_time").Find(&articleList)
	if err != nil {
		return &articleList, errors.New("搜索文章列表错误，数据库搜索错误")
	}
	return &articleList, err
}

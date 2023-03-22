package dao

import (
	"errors"
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	ArticleCommentDao struct {
	}
)

func (ac *ArticleCommentDao) InertComment(articleComment *model.ArticleComment) error {
	orm := util.Orm
	if _, err := orm.InsertOne(articleComment); err != nil {
		return errors.New("评论失败，数据库插入失败")
	}
	return nil
}
func (ac *ArticleCommentDao) GetCommentList(articleComment *model.ArticleComment) (*[]model.ArticleComment, error) {
	orm := util.Orm
	articleCommentList := make([]model.ArticleComment, 0)
	err := orm.Where(&model.ArticleComment{ArticleId: articleComment.ArticleId}).OrderBy("create_time").Find(&articleCommentList)
	if err != nil {
		return &articleCommentList, errors.New("获取评论失败,数据库搜索失败")
	}
	return &articleCommentList, nil
}
func (ac *ArticleCommentDao) SearchCommentById(articleComment *model.ArticleComment) error {
	orm := util.Orm
	if _, err := orm.Get(articleComment); err != nil {
		return errors.New("搜索失败，数据库搜索失败")
	}
	return nil
}

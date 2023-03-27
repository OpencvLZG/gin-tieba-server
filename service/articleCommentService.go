package service

import (
	"ginFlutterBolg/dao"
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
)

type (
	ArticleCommentService struct {
	}
)

func (ac *ArticleCommentService) CreateComment(articleComment *model.ArticleComment) error {
	articleCommentDao := new(dao.ArticleCommentDao)
	articleComment.CommentId = util.GenerateID()
	if err := articleCommentDao.InertComment(articleComment); err != nil {
		return err
	}
	return nil
}
func (ac *ArticleCommentService) GetArticleCommentLimit(articleComment *model.ArticleComment, offset int) (*[]model.ArticleComment, error) {
	articleCommentDao := new(dao.ArticleCommentDao)
	articleCommentList, err := articleCommentDao.GetCommentListOffLim(articleComment, offset)
	if err != nil {
		return articleCommentList, err
	}
	return articleCommentList, nil
}
func (ac *ArticleCommentService) GetArticleComment(articleComment *model.ArticleComment) (*[]model.ArticleComment, error) {
	articleCommentDao := new(dao.ArticleCommentDao)
	articleCommentList, err := articleCommentDao.GetCommentList(articleComment)
	if err != nil {
		return articleCommentList, err
	}
	return articleCommentList, nil
}

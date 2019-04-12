package service

import "go-code-generator/model"

type ArticleService struct {
	Article  *model.Article
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

//获得单一实体
func (article *ArticleService) GetArticleByPrimaryKey() (*model.Article, error) {
	articleRes, err := model.GetArticleByPrimaryKey(article.Article.Id)
	if err != nil {
		return nil, err
	}
	return articleRes, nil
}

//分页获取实体列表
func (article *ArticleService) GetAllArticles() ([]*model.Article, error) {
	articles, err := model.FindArticles(article.PageNum, article.PageSize, article.Article)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

//删除单一实体
func (article *ArticleService) DeleteArticleByPrimaryKey() error {
	return model.DeleteArticleByPrimaryKey(article.Article.Id)
}

//编译单一实体
func (article *ArticleService) EditArticleByPrimaryKey() error {
	return model.EditArticleByPrimaryKey(article.Article.Id, article.Article)
}

//新增单一实体
func (article *ArticleService) AddArticle() error {
	return model.AddArticle(article.Article)
}

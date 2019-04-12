package routes

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-code-generator/model"
	"go-code-generator/service"
)

//获得单一实体
func GetArticleByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	article := &model.Article{Id: id}
	articleService := service.ArticleService{Article: article}
	articleRes, e := articleService.GetArticleByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, articleRes)
}

//分页获取实体列表
func GetAllArticles(c *gin.Context) {
	article := &model.Article{}
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	c.Bind(article)
	articleService := service.ArticleService{
		PageNum:  pageNum,
		PageSize: pageSize,
		Article:  article,
	}
	articleRes, e := articleService.GetAllArticles()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, articleRes)
}

//获得单一实体
func EditArticleByPrimaryKey(c *gin.Context) {
	article := &model.Article{}
	c.Bind(article)
	articleService := service.ArticleService{
		Article: article,
	}
	e := articleService.EditArticleByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

//获得单一实体
func DeleteArticleByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	article := &model.Article{Id: id}
	articleService := service.ArticleService{Article: article}
	e := articleService.DeleteArticleByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

//新增实体
func AddArticle(c *gin.Context) {
	article := &model.Article{}
	c.Bind(article)
	articleService := service.ArticleService{
		Article: article,
	}
	e := articleService.AddArticle()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

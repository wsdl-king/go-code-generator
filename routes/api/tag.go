package routes

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-code-generator/model"
	"go-code-generator/service"
)

//获得单一实体
func GetTagByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tag := &model.Tag{Id: id}
	tagService := service.TagService{Tag: tag}
	tagRes, e := tagService.GetTagByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, tagRes)
}

//分页获取实体列表
func GetAllTags(c *gin.Context) {
	tag := &model.Tag{}
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	c.Bind(tag)
	tagService := service.TagService{
		PageNum:  pageNum,
		PageSize: pageSize,
		Tag:      tag,
	}
	tagRes, e := tagService.GetAllTags()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, tagRes)
}

//获得单一实体
func EditTagByPrimaryKey(c *gin.Context) {
	tag := &model.Tag{}
	c.Bind(tag)
	tagService := service.TagService{
		Tag: tag,
	}
	e := tagService.EditTagByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

//获得单一实体
func DeleteTagByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tag := &model.Tag{Id: id}
	tagService := service.TagService{Tag: tag}
	e := tagService.DeleteTagByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

//新增实体
func AddTag(c *gin.Context) {
	tag := &model.Tag{}
	c.Bind(tag)
	tagService := service.TagService{
		Tag: tag,
	}
	e := tagService.AddTag()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

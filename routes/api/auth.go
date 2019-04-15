package routes

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-code-generator/model"
	"go-code-generator/result"
	"go-code-generator/service"
)

//获得单一实体
func GetAuthByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("key")).MustInt()
	auth := &model.Auth{Id: id}
	authService := service.AuthService{Auth: auth}
	authRes, e := authService.GetAuthByPrimaryKey()
	if e != nil {
		c.JSON(500, result.ErrorResult(authRes, 50000))
	}
	c.JSON(200, result.SuccessResult(authRes))
}

//分页获取实体列表
func GetAllAuths(c *gin.Context) {
	auth := &model.Auth{}
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	c.Bind(auth)
	authService := service.AuthService{
		PageNum:  pageNum,
		PageSize: pageSize,
		Auth:     auth,
	}
	authRes, e := authService.GetAllAuths()
	if e != nil {
		c.JSON(500, result.ErrorResult(authRes, 50000))
	}
	c.JSON(200, result.SuccessResult(authRes))
}

//获得单一实体
func EditAuthByPrimaryKey(c *gin.Context) {
	auth := &model.Auth{}
	c.Bind(auth)
	authService := service.AuthService{
		Auth: auth,
	}
	e := authService.EditAuthByPrimaryKey()
	if e != nil {
		c.JSON(500, result.ErrorResult("", 50000))
	}
	c.JSON(200, result.SuccessResult(""))
}

//获得单一实体
func DeleteAuthByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	auth := &model.Auth{Id: id}
	authService := service.AuthService{Auth: auth}
	e := authService.DeleteAuthByPrimaryKey()
	if e != nil {
		c.JSON(500, result.ErrorResult("", 50000))
	}
	c.JSON(200, result.SuccessResult(""))
}

//新增实体
func AddAuth(c *gin.Context) {
	auth := &model.Auth{}
	c.Bind(auth)
	authService := service.AuthService{
		Auth: auth,
	}
	e := authService.AddAuth()
	if e != nil {
		c.JSON(500, result.ErrorResult("", 50000))
	}
	c.JSON(200, result.SuccessResult(""))
}

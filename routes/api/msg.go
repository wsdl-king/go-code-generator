package routes

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-code-generator/model"
	"go-code-generator/service"
)

//获得单一实体
func GetMsgByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	msg := &model.Msg{Id: id}
	msgService := service.MsgService{Msg: msg}
	msgRes, e := msgService.GetMsgByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, msgRes)
}

//分页获取实体列表
func GetAllMsgs(c *gin.Context) {
	msg := &model.Msg{}
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	c.Bind(msg)
	msgService := service.MsgService{
		PageNum:  pageNum,
		PageSize: pageSize,
		Msg:      msg,
	}
	msgRes, e := msgService.GetAllMsgs()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, msgRes)
}

//获得单一实体
func EditMsgByPrimaryKey(c *gin.Context) {
	msg := &model.Msg{}
	c.Bind(msg)
	msgService := service.MsgService{
		Msg: msg,
	}
	e := msgService.EditMsgByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

//获得单一实体
func DeleteMsgByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	msg := &model.Msg{Id: id}
	msgService := service.MsgService{Msg: msg}
	e := msgService.DeleteMsgByPrimaryKey()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

//新增实体
func AddMsg(c *gin.Context) {
	msg := &model.Msg{}
	c.Bind(msg)
	msgService := service.MsgService{
		Msg: msg,
	}
	e := msgService.AddMsg()
	if e != nil {
		c.JSON(500, "")
	}
	c.JSON(200, "成功")
}

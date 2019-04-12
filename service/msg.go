package service

import "go-code-generator/model"

type MsgService struct {
	Msg      *model.Msg
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

//获得单一实体
func (msg *MsgService) GetMsgByPrimaryKey() (*model.Msg, error) {
	msgRes, err := model.GetMsgByPrimaryKey(msg.Msg.Id)
	if err != nil {
		return nil, err
	}
	return msgRes, nil
}

//分页获取实体列表
func (msg *MsgService) GetAllMsgs() ([]*model.Msg, error) {
	msgs, err := model.FindMsgs(msg.PageNum, msg.PageSize, msg.Msg)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

//删除单一实体
func (msg *MsgService) DeleteMsgByPrimaryKey() error {
	return model.DeleteMsgByPrimaryKey(msg.Msg.Id)
}

//编译单一实体
func (msg *MsgService) EditMsgByPrimaryKey() error {
	return model.EditMsgByPrimaryKey(msg.Msg.Id, msg.Msg)
}

//新增单一实体
func (msg *MsgService) AddMsg() error {
	return model.AddMsg(msg.Msg)
}

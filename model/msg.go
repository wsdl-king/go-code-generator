package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Msg struct {
	Content string `form:"content" json:"content"` // 内容

	Createtime time.Time `form:"createtime" json:"createtime"` //

	Id int `form:"id" json:"id"` //

	ReceiverId int `form:"receiver_id" json:"receiver_id"` // 接收者

	SenderId int `form:"sender_id" json:"sender_id"` // 发送者

	Status int8 `form:"status" json:"status"` //
}

//根据主键得到单一实体
func GetMsgByPrimaryKey(id int) (*Msg, error) {
	var msg Msg
	err := db.Where("id = ?", id).First(&msg).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &msg, nil
}

//根据主键通过条件编辑实体
func EditMsgByPrimaryKey(id int, maps interface{}) error {
	if err := db.Model(&Msg{}).Where("id = ?", id).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

//根据主键删除实体
func DeleteMsgByPrimaryKey(id int) error {
	if err := db.Where("id = ?", id).Delete(&Msg{}).Error; err != nil {
		return err
	}
	return nil
}

//插入实体
func AddMsg(msg *Msg) error {
	if err := db.Create(&msg).Error; err != nil {
		return err
	}
	return nil
}

//根据条件获得分页实体集合
func FindMsgs(pageNum int, pageSize int, maps interface{}) ([]*Msg, error) {
	var (
		msgs []*Msg
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&msgs).Error
	} else {
		err = db.Where(maps).Find(&msgs).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return msgs, nil
}

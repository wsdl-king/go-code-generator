package model

import "github.com/jinzhu/gorm"

type Auth struct {
	Id int `form:"id" json:"id"` //

	Username string `form:"username" json:"username"` // 账号

	Password string `form:"password" json:"password"` // 密码
}

//根据主键得到单一实体
func GetAuthByPrimaryKey(id int) (*Auth, error) {
	var auth Auth
	err := db.Where("id = ?", id).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &auth, nil
}

//根据主键通过条件编辑实体
func EditAuthByPrimaryKey(id int, maps interface{}) error {
	if err := db.Model(&Auth{}).Where("id = ?", id).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

//根据主键删除实体
func DeleteAuthByPrimaryKey(id int) error {
	if err := db.Where("id = ?", id).Delete(&Auth{}).Error; err != nil {
		return err
	}
	return nil
}

//插入实体
func AddAuth(auth *Auth) error {
	if err := db.Create(&auth).Error; err != nil {
		return err
	}
	return nil
}

//根据条件获得分页实体集合
func FindAuths(pageNum int, pageSize int, maps interface{}) ([]*Auth, error) {
	var (
		auths []*Auth
		err   error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&auths).Error
	} else {
		err = db.Where(maps).Find(&auths).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return auths, nil
}

package model

import "github.com/jinzhu/gorm"

type Tag struct {
	CreatedBy string `form:"created_by" json:"created_by"` // 创建人

	CreatedOn int `form:"created_on" json:"created_on"` // 创建时间

	DeletedOn int `form:"deleted_on" json:"deleted_on"` // 删除时间

	Id int `form:"id" json:"id"` //

	ModifiedBy string `form:"modified_by" json:"modified_by"` // 修改人

	ModifiedOn int `form:"modified_on" json:"modified_on"` // 修改时间

	Name string `form:"name" json:"name"` // 标签名称

	State int8 `form:"state" json:"state"` // 状态 0为禁用、1为启用
}

//根据主键得到单一实体
func GetTagByPrimaryKey(id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tag, nil
}

//根据主键通过条件编辑实体
func EditTagByPrimaryKey(id int, maps interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

//根据主键删除实体
func DeleteTagByPrimaryKey(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

//插入实体
func AddTag(tag *Tag) error {
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

//根据条件获得分页实体集合
func FindTags(pageNum int, pageSize int, maps interface{}) ([]*Tag, error) {
	var (
		tags []*Tag
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&tags).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

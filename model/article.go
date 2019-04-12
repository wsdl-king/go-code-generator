package model

import "github.com/jinzhu/gorm"

type Article struct {
	Content string `form:"content" json:"content"` // 内容

	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url"` // 封面图片地址

	CreatedBy string `form:"created_by" json:"created_by"` // 创建人

	CreatedOn int `form:"created_on" json:"created_on"` // 新建时间

	DeletedOn int `form:"deleted_on" json:"deleted_on"` //

	Desc string `form:"desc" json:"desc"` // 简述

	Id int `form:"id" json:"id"` //

	ModifiedBy string `form:"modified_by" json:"modified_by"` // 修改人

	ModifiedOn int `form:"modified_on" json:"modified_on"` // 修改时间

	State int8 `form:"state" json:"state"` // 删除时间

	TagId int `form:"tag_id" json:"tag_id"` // 标签ID

	Title string `form:"title" json:"title"` // 文章标题
}

//根据主键得到单一实体
func GetArticleByPrimaryKey(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

//根据主键通过条件编辑实体
func EditArticleByPrimaryKey(id int, maps interface{}) error {
	if err := db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

//根据主键删除实体
func DeleteArticleByPrimaryKey(id int) error {
	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}

//插入实体
func AddArticle(article *Article) error {
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

//根据条件获得分页实体集合
func FindArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var (
		articles []*Article
		err      error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles).Error
	} else {
		err = db.Where(maps).Find(&articles).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

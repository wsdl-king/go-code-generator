package service

import "go-code-generator/model"

type TagService struct {
	Tag      *model.Tag
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

//获得单一实体
func (tag *TagService) GetTagByPrimaryKey() (*model.Tag, error) {
	tagRes, err := model.GetTagByPrimaryKey(tag.Tag.Id)
	if err != nil {
		return nil, err
	}
	return tagRes, nil
}

//分页获取实体列表
func (tag *TagService) GetAllTags() ([]*model.Tag, error) {
	tags, err := model.FindTags(tag.PageNum, tag.PageSize, tag.Tag)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

//删除单一实体
func (tag *TagService) DeleteTagByPrimaryKey() error {
	return model.DeleteTagByPrimaryKey(tag.Tag.Id)
}

//编译单一实体
func (tag *TagService) EditTagByPrimaryKey() error {
	return model.EditTagByPrimaryKey(tag.Tag.Id, tag.Tag)
}

//新增单一实体
func (tag *TagService) AddTag() error {
	return model.AddTag(tag.Tag)
}

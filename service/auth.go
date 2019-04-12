package service

import "go-code-generator/model"

type AuthService struct {
	Auth     *model.Auth
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

//获得单一实体
func (auth *AuthService) GetAuthByPrimaryKey() (*model.Auth, error) {
	authRes, err := model.GetAuthByPrimaryKey(auth.Auth.Id)
	if err != nil {
		return nil, err
	}
	return authRes, nil
}

//分页获取实体列表
func (auth *AuthService) GetAllAuths() ([]*model.Auth, error) {
	auths, err := model.FindAuths(auth.PageNum, auth.PageSize, auth.Auth)
	if err != nil {
		return nil, err
	}
	return auths, nil
}

//删除单一实体
func (auth *AuthService) DeleteAuthByPrimaryKey() error {
	return model.DeleteAuthByPrimaryKey(auth.Auth.Id)
}

//编译单一实体
func (auth *AuthService) EditAuthByPrimaryKey() error {
	return model.EditAuthByPrimaryKey(auth.Auth.Id, auth.Auth)
}

//新增单一实体
func (auth *AuthService) AddAuth() error {
	return model.AddAuth(auth.Auth)
}

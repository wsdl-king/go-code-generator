package service

import "go-code-generator/model"

type CustomerService struct {
	Customer *model.Customer
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

//获得单一实体
func (customer *CustomerService) GetCustomerByPrimaryKey() (*model.Customer, error) {
	customerRes, err := model.GetCustomerByPrimaryKey(customer.Customer.CustomerId)
	if err != nil {
		return nil, err
	}
	return customerRes, nil
}

//分页获取实体列表
func (customer *CustomerService) GetAllCustomers() ([]*model.Customer, error) {
	customers, err := model.FindCustomers(customer.PageNum, customer.PageSize, customer.Customer)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

//删除单一实体
func (customer *CustomerService) DeleteCustomerByPrimaryKey() error {
	return model.DeleteCustomerByPrimaryKey(customer.Customer.CustomerId)
}

//编译单一实体
func (customer *CustomerService) EditCustomerByPrimaryKey() error {
	return model.EditCustomerByPrimaryKey(customer.Customer.CustomerId, customer.Customer)
}

//新增单一实体
func (customer *CustomerService) AddCustomer() error {
	return model.AddCustomer(customer.Customer)
}

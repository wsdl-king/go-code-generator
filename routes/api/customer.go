package routes

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-code-generator/model"
	"go-code-generator/result"
	"go-code-generator/service"
)

//获得单一实体
func GetCustomerByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("key")).MustInt()
	customer := &model.Customer{CustomerId: id}
	customerService := service.CustomerService{Customer: customer}
	customerRes, e := customerService.GetCustomerByPrimaryKey()
	if e != nil {
		c.JSON(500, result.ErrorResult(customerRes, 50000))
	}
	c.JSON(200, result.SuccessResult(customerRes))
}

//分页获取实体列表
func GetAllCustomers(c *gin.Context) {
	customer := &model.Customer{}
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	c.Bind(customer)
	customerService := service.CustomerService{
		PageNum:  pageNum,
		PageSize: pageSize,
		Customer: customer,
	}
	customerRes, e := customerService.GetAllCustomers()
	if e != nil {
		c.JSON(500, result.ErrorResult(customerRes, 50000))
	}
	c.JSON(200, result.SuccessResult(customerRes))
}

//获得单一实体
func EditCustomerByPrimaryKey(c *gin.Context) {
	customer := &model.Customer{}
	c.Bind(customer)
	customerService := service.CustomerService{
		Customer: customer,
	}
	e := customerService.EditCustomerByPrimaryKey()
	if e != nil {
		c.JSON(500, result.ErrorResult("", 50000))
	}
	c.JSON(200, result.SuccessResult(""))
}

//获得单一实体
func DeleteCustomerByPrimaryKey(c *gin.Context) {
	id := com.StrTo(c.Param("customerId")).MustInt()
	customer := &model.Customer{CustomerId: id}
	customerService := service.CustomerService{Customer: customer}
	e := customerService.DeleteCustomerByPrimaryKey()
	if e != nil {
		c.JSON(500, result.ErrorResult("", 50000))
	}
	c.JSON(200, result.SuccessResult(""))
}

//新增实体
func AddCustomer(c *gin.Context) {
	customer := &model.Customer{}
	c.Bind(customer)
	customerService := service.CustomerService{
		Customer: customer,
	}
	e := customerService.AddCustomer()
	if e != nil {
		c.JSON(500, result.ErrorResult("", 50000))
	}
	c.JSON(200, result.SuccessResult(""))
}

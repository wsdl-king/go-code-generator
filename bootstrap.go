package main

import (
	"github.com/gin-gonic/gin"
	"go-code-generator/routes/api"
	"net/http"
	"time"
)

func main() {
	engine := gin.New()
	//自定义
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           engine,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	//根据主键获得单一商户表

	engine.GET("/getCustomer/:customerId", routes.GetCustomerByPrimaryKey)

	//获取分页商户表

	engine.POST("/findCustomerList", routes.GetAllCustomers)

	//修改商户表

	engine.POST("/updateCustomer", routes.EditCustomerByPrimaryKey)

	//删除商户表

	engine.DELETE("/deleteCustomer", routes.DeleteCustomerByPrimaryKey)

	//新增商户表

	engine.POST("/AddCustomer", routes.AddCustomer)

	srv.ListenAndServe()
}

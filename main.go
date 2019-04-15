package main

import (
	"go-code-generator/conf"
	"go-code-generator/generate"
)

func main() {
	// 生成 go-gin 启动类
	//generate.BootStrap()
	// 读取数据库,生成route_api
	env := conf.GetEnv()
	generate.GenerateRoute(env.DBName, "")
	//db.InitDB()
	//maps := make(map[string]interface{})
	//maps["name"] = "保定"
	//customer := &model.Customer{Name: "邯郸",Lon:132.45,Lat:45.21}
	//customerService :=  service.CustomerService{Customer:customer}
	//customerService.AddCustomer()
	//customers, _:= customerService.GetAll()
	//get := customerService.EditCustomerByCustomerId()
	//fmt.Println(customers)
	//fmt.Println(get)
}

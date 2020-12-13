# go-code-generator [![rcard](https://goreportcard.com/badge/github.com/wsdl-king/go-code-generator)](https://goreportcard.com/report/github.com/wsdl-king/go-code-generator) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/wsdl-king/go-code-generator) 
go写的自动生成代码工具,一键生成go-gin+gorm的增删改查,只需要配置一下数据库就可以,先起草一个初始化版本.


## 安装
```
$ go get github.com/wsdl-king/go-code-generator
```
## 如何运行

### 必须

- Mysql
- GoLang1.6+

### 依赖
```
go get  https://github.com/jinzhu/gorm
go get  https://github.com/gin-gonic/gin
```

### 配置

你需要修改 `config.yml` 配置文件

```
 #开发环境
 development:
   database:
     host: 127.0.0.1
     user: root
     pwd: 123456
     port: 3306
     #数据库名称
     db_name: test
     #数据库表名称
     table_name: customer
     #DriverName
     driver_name: mysql
``` 
### 运行
-step:
```
$ cd $GOPATH/src/go-code-generator

$ go run main.go 
```
-step2:
```
$ cd $GOPATH/src/go-code-generator

$ go run bootstrap.go 
```

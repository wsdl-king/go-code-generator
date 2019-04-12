package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	Setup()
	fmt.Println("加载成功")
}

func Setup() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1",
		"3306", "test"))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	// 打印sql
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

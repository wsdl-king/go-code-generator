package generate

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-code-generator/db"
)

// 生成增删改查的route
func GenerateRoute(dbName, tableName string) {
	db := db.GetDB()
	generateFromDataBase(db, dbName, tableName)
}

//读取数据库名/表名
func generateFromDataBase(db *sqlx.DB, dbName, tableName string) {
	AllTable := &[]ALLTABLE{}
	if dbName == "" && tableName == "" {
		return
	}
	var sql string
	if tableName == "" {
		sql = "SELECT TABLE_NAME,TABLE_COMMENT from information_schema.`TABLES` where " +
			"TABLE_SCHEMA = '" + dbName + "'"
	} else {
		sql = "SELECT TABLE_NAME,TABLE_COMMENT from information_schema.`TABLES` where " +
			"TABLE_SCHEMA = '" + dbName + "'" + " and TABLE_NAME='" + tableName + "' "
	}
	e := db.Select(AllTable, sql)
	if e != nil {
		fmt.Println(e)
		return
	}
	for _, table := range *AllTable {
		//生成model/dao
		generateStart(db, table.TABLE_NAME)
	}
	buildRoutes(AllTable)
}

//生成model/dao
func generateStart(db *sqlx.DB, tableName string) {
	//拿到 tableName 进行渲染 生成model
	start(db, tableName, "test")
}

// 生成go-gin启动类

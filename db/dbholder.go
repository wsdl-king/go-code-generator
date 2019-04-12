package db

import (
	"github.com/jmoiron/sqlx"
	_ "gopkg.in/go-sql-driver/mysql.v1"
	"strings"
)

//数据库配置
const (
	userName            = "root"
	password            = "123456"
	ip                  = "127.0.0.1"
	port                = "3306"
	dbName              = "test"
	driverName          = "mysql"
	default_db_max_open = 32
	default_db_max_idle = 2
)

var DB *sqlx.DB

func init() {
	InitDB()
}
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, _ := sqlx.Open(driverName, path)
	db.SetMaxOpenConns(default_db_max_open)
	db.SetMaxIdleConns(default_db_max_idle)
	DB = db
}
func GetDB() *sqlx.DB {
	return DB
}

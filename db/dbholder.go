package db

import (
	"github.com/jmoiron/sqlx"
	"go-code-generator/conf"
	_ "gopkg.in/go-sql-driver/mysql.v1"
	"strings"
)

//数据库配置--只连接一次我就不拿到配置文件了
const (
	defaultDbMaxOpen = 32
	defaultDbMaxIdle = 2
)

var DB *sqlx.DB
var config = conf.GetEnv()

func init() {
	InitDB()
}
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{config.DbUser, ":", config.DbPwd, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DBName, "?charset=utf8"}, "")
	db, _ := sqlx.Open(config.DriverName, path)
	db.SetMaxOpenConns(defaultDbMaxOpen)
	db.SetMaxIdleConns(defaultDbMaxIdle)
	DB = db
}
func GetDB() *sqlx.DB {
	return DB
}

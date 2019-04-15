package conf

import (
	"io/ioutil"
)

type DBConfig struct {
	DbUser     string
	DbPwd      string
	DbHost     string
	DbPort     string
	DBName     string
	TableName  string
	DriverName string
}

type Env struct {
	//环境 -- 开发还是生产
	Profile string
	//数据库配置类
	DBConfig
	//一个结构切片
}

func init() {
	NewConfig("")
}

var defaultPath = "./config.yml"
var env *Env

func NewConfig(path string) {
	if path == "" {
		path = defaultPath
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	yamlString := string(file)
	cfg, err := ParseYaml(yamlString)
	env = &Env{
		Profile: cfg.getConfigString("development"),
		DBConfig: DBConfig{
			DbUser:     cfg.getConfigString("development.database.user"),
			DbPwd:      cfg.getConfigString("development.database.pwd"),
			DbHost:     cfg.getConfigString("development.database.host"),
			DbPort:     cfg.getConfigString("development.database.port"),
			DBName:     cfg.getConfigString("development.database.db_name"),
			TableName:  cfg.getConfigString("development.database.table_name"),
			DriverName: cfg.getConfigString("development.database.driver_name"),
		},
	}
}

func (cfg *Config) getConfigString(key string) string {
	s, _ := cfg.String(key)
	return s
}

func (cfg *Config) getConfigInt(key string) int {
	s, _ := cfg.Int(key)
	return s
}
func GetEnv() *Env {
	return env
}

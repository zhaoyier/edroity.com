package conf

import (
	"edroity.com/server/common/conf/json"
)

const (
	ConfFile = "./conf/server.json"
)

var ServerConf DBConf

// mysql的數據配置結構體
type MysqlServer struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
}

// redis的數據配置結構體
type RedisServer struct {
	Network   string
	Address   string
	MaxIdle   int
	MaxActive int
}

type DBConf struct {
	MysqlConf MysqlServer `json:"mysql"`
	RedisConf RedisServer `json:"redis"`
}

func init() {
	ParseServerConf()
}

//将server 的json config文件转换成json对象
func ParseServerConf() {
	json.Parse(ConfFile, &ServerConf)
}

package main

import (
	"edroity.com/server/pond/model/static"
	"edroity.com/server/common/log"

	"edroity.com/server/common/db/mongo"
	"runtime"
	"net"
	"edroity.com/server/common/conf/json"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var config struct{
	Address string `json:"address"`
}

func main()  {
	defer log.Shutdown()
	defer mongo.Shutdown()

	if err := json.Parse("./config/init.json", &config); err != nil {
		log.Error("[main] 解析配置文件失败，Err:", err)
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	//启动服务
	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Error("[main] 监听端口异常, Err:", err)
		os.Exit(1)
	}
	server := grpc.NewServer()

	_ = static.GetItem("270046")
	log.Debug("[main] 服务器启动.")
}

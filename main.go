package main

import (
	"ginFlutterBolg/router"
	"ginFlutterBolg/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + "/config.json"
	config := util.ReadJson(path)
	print(config)
	////实例化数据库
	_, err = util.GetOrmEngine()
	if err != nil {
		panic(err)
	}
	//实例化路由
	r := gin.Default()
	//注册路由
	router.InitRouter(r)
	r.Run(config.Server.Address + ":" + config.Server.Port)
}

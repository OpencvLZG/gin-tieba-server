package main

import (
	"flag"
	"ginFlutterBolg/router"
	"ginFlutterBolg/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {

	defer func() {
		log.Println("exit server program")
	}()
	serverType := flag.String("serverType", "http", "server type,example: http,https,All")
	cert := flag.String("cert", "./cert/example/server.cert", "the server cert")
	key := flag.String("key", "./cert/example/server.key", "the server key")
	flag.Parse()
	util.LogInit()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + "/config.json"
	log.Println("loading Config,config path" + path)
	config := util.ReadJson(path)
	log.Printf("server config,%v\n", config)
	////实例化数据库
	_, err = util.GetOrmEngine()
	if err != nil {
		panic(err)
	}
	//实例化路由
	r := gin.Default()
	//注册路由
	router.InitRouter(r)
	if *serverType == "https" {
		err = r.RunTLS(config.Server.Address+":"+config.Server.Port, *cert, *key)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = r.Run(config.Server.Address + ":" + config.Server.Port)
		if err != nil {
			log.Fatal(err)
		}
	}

}

package main

import (
	"ginFlutterBolg/model"
	"ginFlutterBolg/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"xorm.io/xorm"
)

func CreateDb() {
	config := util.Cfg.Mysql
	engine, err := xorm.NewEngine(
		"mysql",
		config.User+":"+config.Password+"@("+config.Ip+":"+config.Port+")/")
	if err != nil {
		log.Fatal("Connect Mysql Error", err)
	}

	engine.ShowSQL(true)
	_, err = engine.Exec("CREATE DATABASE " + config.DatabaseName + " DEFAULT CHARACTER SET utf8")

	if err != nil {
		log.Fatal("Create Db Error", err)
	}
}

func InitDb() {
	config := util.Cfg.Mysql
	engine, err := xorm.NewEngine(
		"mysql",
		config.User+":"+config.Password+"@("+config.Ip+":"+config.Port+")/"+config.DatabaseName+"?charset="+config.Charset)
	if err != nil {
		log.Fatal("Connect Mysql Database Error", err)
	}
	engine.Sync2(new(model.User))
	engine.Sync2(new(model.UserFollow))
	engine.Sync2(new(model.Article))
	engine.Sync2(new(model.ArticleComment))
	engine.Sync2(new(model.Belong))
	engine.Sync2(new(model.BelongFollower))
	log.Println("Init Database Successful")
}
func main() {
	//获取配置文件
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + "/config.json"
	println(path)
	_ = util.ReadJson(path)
	CreateDb()
	InitDb()

}

# ginTiebaServer

#### Description
使用golang-gin框架完成的服务后端

#### Software Architecture
ginTiebaServer is base on golang-gin backend develop,a TiebaServer.

#### Installation

1. Install Golang [Google]([Download and install - The Go Programming Language (google.cn)](https://golang.google.cn/doc/install) "go language")

2. Turn On Terminal Cd This Project(You will find the go.mod)

3. install depend

   ```go
   go mod tidy
   ```

4. Create database

   ```mysql
   create database tieba;
   ```

5. Modify the config.json

   ```json
   {
     "app": {
       "appName": "GinServer"
     },
     "server": {
       "address": "127.0.0.1",	
       "port": "80"
     },
     "mysql": {
       "user": "root",
       "password": "123456",
       "ip": "127.0.0.1",
       "port": "3306",
       "databaseName": "tieba",
       "Charset": "utf8"
     }
   }
   ```



6. Remove annotation from "util/xorm.go"

   ```go
   //engine.Sync2(new(model.User))
   //engine.Sync2(new(model.UserFollow))
   ////engine.Sync2(new(model.Article))
   ////engine.Sync2(new(model.ArticleComment))
   //engine.Sync2(new(model.Belong))
   //engine.Sync2(new(model.BelongFollower))
   // This Code will be sync database
   ```

7. Test the ServerAddress:ServerPort

#### Instructions

1.  x
2.  xxxx
3.  xxxx

#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


#### Gitee Feature

1.  

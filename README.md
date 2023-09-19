
<div align=center></br></br></br>

<center> <img src="https://thirdqq.qlogo.cn/g?b=sdk&k=iaNcdgTAPWOS0JJseiafW1Dw&kti=ZIsqGgAAAAI&s=40&t=1638804590" style="zoom:300%;" /></center>

#  <center>ginTiebaServer</center>

###### <center>Introduce</center>

###### 							
![golang](https://img.shields.io/badge/golang-blue?logo=go) ![gin](https://img.shields.io/badge/gin-blue?logo=backend)  ![open sourse (shields.io)](https://img.shields.io/badge/open%20sourse-darkgreen?logo=opensourceinitiative)   ![github (shields.io)](https://img.shields.io/badge/github-grey?logo=github) ![gitee (shields.io)](https://img.shields.io/badge/gitee-orange?logo=gitee) ![git (shields.io)](https://img.shields.io/badge/git-lightblue?logo=git) ![Mit: license (shields.io)](https://img.shields.io/badge/Mit-license-blue?logo=bookstack) ![img](https://komarev.com/ghpvc/?username=cilang-ginTiebaServer&&style=flat-square)

<center>a tieba backend server</center>

</div>

#### Description
based on golang implement
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

#### Api Document


[Postman Document Online](https://documenter.getpostman.com/view/20660781/2s9YC5wrgi)


### Donate

Thank ur supprt

| <center>微信</center>                            | <center>支付宝</center>                       |
| ------------------------------------------------ | --------------------------------------------- |
| <center>![image](./gitImage/weixin.png)</center> | <center>![image](./gitImage/zfb.jpg)</center> |
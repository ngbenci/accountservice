package main

import (
	"goblog/accountservice/dbclient"
	"goblog/accountservice/service"
)

var appName = "accountservice"
var port = "8080"

func main() {

	service.DbClient = &dbclient.BoltClient{}

	//连接数据库
	service.DbClient.OpenBoltDb()

	//初始化桶和账号数据
	service.DbClient.Seed()

	//启动Http服务
	service.StartWebService(port, appName)
}

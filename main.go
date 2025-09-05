package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	//创建数据库
	//sql:CREATE DATABASE GoList
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run()
}

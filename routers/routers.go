package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改第id个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除第id个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleTodo)
	}
	return r
}

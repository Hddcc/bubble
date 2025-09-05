package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//前端页面填写待办事项，点击“+”会发送请求到这里
	//1、从请求中获取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//2、存入数据库并返回响应
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, todo)
	}
}

func GetTodoList(c *gin.Context) {
	//查看todo这个表里的所有数据
	todolist, err := models.GetTodoList()
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, todolist)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id读取错误,可能不存在"})
		return
	}
	todo, err := models.GetTodo(id)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id读取错误,可能不存在"})
		return
	}
	if err := models.DeleTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

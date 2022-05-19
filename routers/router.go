package routers

import (
	"bubble/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", api.IndexHandler)
	v1Group := r.Group("v1")
	{ //代办事项
		//添加
		v1Group.POST("/todo", api.AddTodo)
		//查看所有的待办事项
		v1Group.GET("/todo", api.GetAllTodos)
		//修改某一待办事项
		v1Group.PUT("/todo/:id", api.UpdateTodo)
		//删除某一待办事项
		v1Group.DELETE("/todo/:id", api.DeleteTodo)
	}
	return r
}

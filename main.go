package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mikuLuka/todolist/controller"
	"github.com/mikuLuka/todolist/databaseAccessObject"
	"github.com/mikuLuka/todolist/models"
)

func main() {
	//创建，连接数据库
	err := databaseAccessObject.InitMySQL()
	if err != nil {
		panic(err)
	}
	//模型绑定
	databaseAccessObject.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()
	r.Static("/static", "static") //找静态资源
	r.LoadHTMLGlob("templates/*") //找模板文件
	r.GET("/", controller.IndexHandler)
	//增删改查
	v1Group := r.Group("v1")
	{
		//待办事项，添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("todo/:id", controller.DelateTodo)

	}

	r.Run(":8080")

}

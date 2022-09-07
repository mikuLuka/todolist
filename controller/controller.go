package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mikuLuka/todolist/databaseAccessObject"
	"github.com/mikuLuka/todolist/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	//3.返回响应
	err := databaseAccessObject.DB.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	//查询todos表里所有数据
	var todoList []models.Todo
	err := databaseAccessObject.DB.Find(&todoList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
	}
	var todo models.Todo
	err := databaseAccessObject.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err := databaseAccessObject.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DelateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
	}
	err := databaseAccessObject.DB.Where("id=?", id).Delete(models.Todo{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}

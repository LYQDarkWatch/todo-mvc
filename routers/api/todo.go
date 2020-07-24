package api

import (
	"net/http"
	"todo-mvc/models"
	"todo-mvc/pkg/error"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// GetTodosByUserName 用户获取 todos 信息，可选参数 completion 筛选是否完成
func GetTodosByUserName(c *gin.Context) {
	username := Username
	data := make(map[string]interface{})
	completion, ok := c.GetQuery("completion")
	if !ok {
		data["todos"] = models.GetTodoByUserName(username)
	} else {
		data["todos"] = models.GetTodoByTodoStatus(username, completion)
	}
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

var code int
var Username string

// CreateTodo 创建新的 todo
func CreateTodo(c *gin.Context) {
	todo := models.Todo{}
	c.BindJSON(&todo)
	todo.Owner = Username
	todo.ID = bson.NewObjectId()
	todo.Completion = "false"
	if models.CreateTodo(todo) == false {
		code = error.INVALID_PARAMS
	} else {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

// DeleteTodoByID 删除指定 todo
func DeleteTodoByID(c *gin.Context) {
	var code int
	objectID := c.Query("objectID")
	if models.DeleteTodo(objectID) == false {
		code = error.ERROR_TODO_DELETE_ERROE
	} else {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

// DeleteTodoByOwner 用户清空已完成 todo
func DeleteTodoByOwner(c *gin.Context) {
	var code int
	userName := Username
	if models.DeleteTodosByUser(userName) == false {
		code = error.ERROR_TODO_DELETE_BYNAME
	} else {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

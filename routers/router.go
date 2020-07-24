package routers

import (
	"todo-mvc/middleware/cors"
	"todo-mvc/middleware/jwt"
	"todo-mvc/pkg/setting"
	"todo-mvc/routers/api"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化 gin 路由
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Cors())

	gin.SetMode(setting.RunMode)
	router.POST("/user/login", api.CheckUser)
	router.POST("/user/register", api.Register)
	apiv1 := router.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/todos", api.GetTodosByUserName)

		apiv1.POST("/todos", api.CreateTodo)

		apiv1.DELETE("/todos", api.DeleteTodoByID)

		apiv1.DELETE("/todos/batchDelete", api.DeleteTodoByOwner)

		apiv1.PUT("/todos/complete", api.CompleteAllTodos)
	}
	return router
}

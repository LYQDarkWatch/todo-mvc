package jwt

import (
	"net/http"
	"time"
	"todo-mvc/pkg/error"
	"todo-mvc/pkg/util"
	"todo-mvc/routers/api"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = error.SUCCESS
		token := c.Request.Header.Get("Token")
		if token == "" {
			code = error.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			api.Username = claims.Username
			if err != nil {
				code = error.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = error.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != error.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  error.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

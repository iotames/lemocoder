package prepare

import (
	"lemocoder/database"
	"lemocoder/webserver/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Accept, Token, Auth-Token, X-Requested-With")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Header("Access-Control-Allow-Origin", "*")
		// 请求前
		c.Next()
		// 请求后
	}
}

func HandlerJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Auth-Token")
		if token == "" {
			c.JSON(200, handler.ResponseFail("缺少token参数，无权访问", 401))
			c.Abort()
			return
		}
		u := database.User{}
		user, err := u.GetUserByJwt(token)
		if err != nil {
			c.JSON(200, handler.ResponseFail("鉴权错误:"+err.Error(), 401))
			c.Abort()
			return
		}
		c.Set("user", user)
	}
}

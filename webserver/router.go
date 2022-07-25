package webserver

import (
	"fmt"
	"lemocoder/database"
	"lemocoder/webserver/handler"
	"lemocoder/webserver/prepare"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func checkUserByJWT() gin.HandlerFunc {
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

func localhostNetwork() gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteAddr := c.Request.RemoteAddr
		if strings.Contains(remoteAddr, "127.0.0.1") {
			return
		}
		if strings.Contains(remoteAddr, "::1") {
			return
		}
		c.JSON(200, handler.ResponseFail("仅限内网访问", 401))
		c.Abort()
	}
}

func setRouters(g *gin.Engine) {
	g.MaxMultipartMemory = 8 << 20                     // 8M
	g.StaticFS("/client", http.Dir("resource/client")) // OK
	// g.Static("/uploads", "runtime/uploads")            // OK
	g.Use(prepare.HandlerCORS())
	g.GET("/debug", debug)

	api := g.Group("/api")
	api.POST("/public/login", handler.Login)
	api.GET("/client/config", handler.GetClientConfig)
	api.POST("/client/init", handler.ClientInit)

	local := api.Group("/local")
	local.Use(localhostNetwork())
	local.POST("/upload", handler.UploadFile)
	local.POST("/excelspider", handler.ExcelSpider)

	user := api.Group("/user")
	user.Use(checkUserByJWT())
	user.GET("/info", handler.GetUserInfo)
	user.POST("/logout", handler.Logout)
	user.GET("menu", handler.GetClientMenu)
}

func debug(c *gin.Context) {
	log.Println("----Request--", c.RemoteIP(), c.Request.RemoteAddr)
	c.String(200, "RemoteIP():%s--Request.Host:%s--", c.RemoteIP(), c.Request.Host)
	c.String(200, fmt.Sprintf("---Success--Debug-----\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_DRIVER\"):(%v)---", os.Getenv("DB_HOST"), os.Getenv("DB_DRIVER")))
}

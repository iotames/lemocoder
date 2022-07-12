package webserver

import (
	"fmt"
	"lemocoder/webserver/handler"
	"lemocoder/webserver/prepare"
	"os"

	"github.com/gin-gonic/gin"
)

func setRouters(g *gin.Engine) {
	g.Use(prepare.HandlerCORS())
	g.GET("/gdebug", gdebug)
	g.GET("/api/currentUser", handler.GetUserInfo)
	g.GET("/api/client/config", handler.GetClientConfig)
	g.POST("/api/user/logout", handler.Logout)
}

func gdebug(c *gin.Context) {
	c.String(200, fmt.Sprintf("---Success--Debug-----\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_TYPE\"):(%v)---", os.Getenv("DB_HOST"), os.Getenv("DB_TYPE")))
}

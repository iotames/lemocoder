package webserver

import (
	"fmt"
	"lemocoder/webserver/handler"
	"lemocoder/webserver/prepare"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setRouters(g *gin.Engine) {
	g.StaticFS("/client", http.Dir("resource/client")) // OK
	// g.Static("/client", "resource/client") // OK
	g.Use(prepare.HandlerCORS())
	g.GET("/debug", debug)
	g.POST("/api/public/upload", uploadFile)
	g.POST("/api/public/login", handler.Login)
	g.GET("/api/client/config", handler.GetClientConfig)
	g.GET("/api/user/info", handler.GetUserInfo)
	g.POST("/api/user/logout", handler.Logout)
	g.GET("/api/user/menu", handler.GetClientMenu)
}

func debug(c *gin.Context) {
	c.String(200, fmt.Sprintf("---Success--Debug-----\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_DRIVER\"):(%v)---", os.Getenv("DB_HOST"), os.Getenv("DB_DRIVER")))
}

func uploadFile(c *gin.Context) {
	c.String(200, "helooooooo")
}

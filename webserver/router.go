package webserver

import (
	"fmt"
	"lemocoder/webserver/handler"
	"lemocoder/webserver/prepare"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setRouters(g *gin.Engine) {
	g.StaticFS("/client", http.Dir("resource/client")) // OK
	// g.Static("/uploads", "runtime/uploads")            // OK
	g.Use(prepare.HandlerCORS())
	g.GET("/debug", debug)
	g.MaxMultipartMemory = 8 << 20 // 8M

	g.POST("/api/local/excelspider", handler.ExcelSpider)
	g.POST("/api/local/upload", handler.UploadFile)

	g.POST("/api/public/login", handler.Login)
	g.GET("/api/client/config", handler.GetClientConfig)
	g.GET("/api/user/info", handler.GetUserInfo)
	g.POST("/api/user/logout", handler.Logout)
	g.GET("/api/user/menu", handler.GetClientMenu)
}

func debug(c *gin.Context) {
	log.Println("----Request--", c.RemoteIP(), c.Request.RemoteAddr)
	c.String(200, "RemoteIP():%s--Request.Host:%s--", c.RemoteIP(), c.Request.Host)
	c.String(200, fmt.Sprintf("---Success--Debug-----\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_DRIVER\"):(%v)---", os.Getenv("DB_HOST"), os.Getenv("DB_DRIVER")))
}

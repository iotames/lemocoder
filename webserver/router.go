package webserver

import (
	"fmt"
	"lemocoder/webserver/handler"
	"lemocoder/webserver/prepare"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

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
	api.GET("/client/createcode", handler.CreateCode)
	api.POST("/public/login", handler.Login)
	api.GET("/client/config", handler.GetClientConfig)
	api.POST("/client/init", handler.ClientInit)
	api.GET("/table/demodata", getTableDataDemo)

	local := api.Group("/local")
	local.Use(localhostNetwork())
	local.POST("/upload", handler.UploadFile)
	local.POST("/excelspider", handler.ExcelSpider)
	local.POST("/stop", handler.StopApp)

	auth := api.Group("")
	auth.Use(prepare.HandlerJWT())

	auth.GET("/os/getstatus", handler.GetOsStatus)

	auth.POST("/coder/project/rebuild", handler.RebuildProject)
	auth.POST("/coder/page/create", handler.CreateWebPage)
	auth.POST("/coder/page/delete", handler.DeleteWebPage)
	auth.POST("/coder/page/update", handler.UpdateWebPage)
	auth.GET("/coder/page/get", handler.GetWebPage)
	auth.GET("/coder/pages", handler.GetWebPagesList)
	auth.GET("/coder/table/get", handler.GetTable)
	auth.POST("/coder/table/create", handler.CreateTable)
	auth.POST("/coder/table/update", handler.UpdateTable)
	auth.POST("/coder/table/createcode", handler.CreateTablePageCode)

	// ADD AUTO CODE
	addRoutes(auth)

	user := auth.Group("/user")
	user.GET("/info", handler.GetUserInfo)
	user.POST("/logout", handler.Logout)
	user.GET("menu", handler.GetClientMenu)
}

func debug(c *gin.Context) {
	log.Println("----Request--", c.RemoteIP(), c.Request.RemoteAddr)
	c.String(200, "RemoteIP():%s--Request.Host:%s--", c.RemoteIP(), c.Request.Host)
	c.String(200, fmt.Sprintf("---Success--Debug-----\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_DRIVER\"):(%v)---", os.Getenv("DB_HOST"), os.Getenv("DB_DRIVER")))
}

func getTableDataDemo(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	title := c.Query("title")
	sort := c.Query("sort")
	log.Printf("----page:%s--limit:%s---title:%s---sort:%s", page, limit, title, sort)
	// ----page:1--limit:10---title:您好世界---sort:{"created_at":"ascend"} ---sort:{"created_at":"descend"}
	resp := `{"Data":{"Items":[
		{"id":624748504,"number":6689,"title":"🐛 [BUG]yarn install命令 antd2.4.5会报错","labels":[{"name":"bug","color":"error"}],"state":"open","locked":false,"comments":1,"created_at":"2020-05-26T09:42:56Z","updated_at":"2020-05-26T10:03:02Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624691229,"number":6688,"title":"🐛 [BUG]无法创建工程npm create umi","labels":[{"name":"bug","color":"success"}],"state":"closed","locked":false,"comments":0,"created_at":"2020-05-26T08:19:22Z","updated_at":"2020-05-26T08:19:22Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624674790,"number":6685,"title":"🧐 [问题] build 后还存在 es6 的代码（Umi@2.13.13）","labels":[{"name":"question","color":"success"}],"state":"open","locked":false,"comments":0,"created_at":"2020-05-26T07:54:25Z","updated_at":"2020-05-26T07:54:25Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624620220,"number":6683,"title":"2.3.1版本如何在业务页面修改头部状态","labels":[{"name":"question","color":"success"}],"state":"processing","locked":false,"comments":2,"created_at":"2020-05-26T05:58:24Z","updated_at":"2020-05-26T07:17:39Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624592471,"number":6682,"title":"hideChildrenInMenu设置后，子路由找不到了","labels":[{"name":"bug","color":"error"}],"state":"open","locked":false,"comments":2,"created_at":"2020-05-26T04:25:59Z","updated_at":"2020-05-26T08:00:51Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624748503,"number":6689,"title":"🐛 [BUG]yarn install命令 antd2.4.5会报错","labels":[{"name":"bug","color":"error"}],"state":"open","locked":false,"comments":1,"created_at":"2020-05-26T09:42:56Z","updated_at":"2020-05-26T10:03:02Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624691228,"number":6688,"title":"🐛 [BUG]无法创建工程npm create umi","labels":[{"name":"bug","color":"error"}],"state":"open","locked":false,"comments":0,"created_at":"2020-05-26T08:19:22Z","updated_at":"2020-05-26T08:19:22Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624674789,"number":6685,"title":"🧐 [问题] build 后还存在 es6 的代码（Umi@2.13.13）","labels":[{"name":"question","color":"success"}],"state":"open","locked":false,"comments":0,"created_at":"2020-05-26T07:54:25Z","updated_at":"2020-05-26T07:54:25Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624620219,"number":6683,"title":"2.3.1版本如何在业务页面修改头部状态","labels":[{"name":"question","color":"success"}],"state":"open","locked":false,"comments":2,"created_at":"2020-05-26T05:58:24Z","updated_at":"2020-05-26T07:17:39Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"},
		{"id":624592470,"number":6682,"title":"hideChildrenInMenu设置后，子路由找不到了","labels":[{"name":"bug","color":"error"}],"state":"open","locked":false,"comments":2,"created_at":"2020-05-26T04:25:59Z","updated_at":"2020-05-26T08:00:51Z","closed_at":null,"author_association":"NONE","user":"chenshuai2144","avatar":"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"}
		],"Page":1,"Total":30},"Msg":"SUCCESS", "Code":200}`
	c.String(200, resp)
}

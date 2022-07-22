package handler

import (
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/generator"
	"lemocoder/util"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ClientConfig struct {
	IsLocked                                                                                   bool
	Title, Logo, DbDriver, DbHost, DbName, DbPassword, DbUsername, LoginAccount, LoginPassword string
	DbNodeId, DbPort, WebServerPort                                                            int
}

func GetClientConfig(c *gin.Context) {
	d := config.GetDatabase()
	s := config.GetWebServer()
	a := config.GetApp()
	conf := ClientConfig{
		IsLocked:      util.IsPathExists("app.lock"),
		Title:         a.Title,
		Logo:          a.Logo,
		DbDriver:      d.Driver,
		DbHost:        d.Host,
		DbName:        d.Name,
		DbUsername:    d.Username,
		DbPassword:    d.Password,
		DbNodeId:      d.NodeID,
		DbPort:        d.Port,
		WebServerPort: s.Port,
	}
	c.JSON(http.StatusOK, Response(conf, "success", 200))
}

func ClientInit(c *gin.Context) {
	conf := new(ClientConfig)
	c.Bind(conf)
	if conf.Logo == "" {
		a := config.GetApp()
		conf.Logo = a.Logo
	}
	lockFile := "app.lock"
	if util.IsPathExists(lockFile) {
		c.JSON(http.StatusOK, ResponseFail("请不要重复初始化", 400))
		return
	}
	if len(conf.LoginPassword) < 6 {
		c.JSON(http.StatusOK, ResponseFail("密码长度过短", 400))
		return
	}
	u := database.User{Account: conf.LoginAccount}
	u, err := u.Register(conf.LoginPassword)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("用户注册失败: "+err.Error(), 500))
		return
	}
	f, err := os.OpenFile(config.EnvFilepath, os.O_RDWR, 0644)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("open .env File Error: "+err.Error(), 500))
		return
	}
	defer f.Close()
	err = generator.GetContentByTpl(config.TplFilepath+"/env.tpl", f, conf)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("create .env File Error: "+err.Error(), 500))
		return
	}
	f, _ = os.Create(lockFile)
	f.Close()
	c.JSON(http.StatusOK, ResponseOk("success"))
}

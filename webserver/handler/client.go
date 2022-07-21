package handler

import (
	"lemocoder/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ClientConfig struct {
	Title, Logo, DbDriver, DbHost, DbName, DbPassword, DbUsername string
	DbNodeId, DbPort, WebPort                                     int
}

func GetClientConfig(c *gin.Context) {
	d := config.GetDatabase()
	s := config.GetWebServer()
	conf := ClientConfig{
		Title:      os.Getenv("APP_TITLE"),
		Logo:       os.Getenv("APP_LOGO"),
		DbDriver:   d.Driver,
		DbHost:     d.Host,
		DbName:     d.Name,
		DbUsername: d.Username,
		DbPassword: d.Password,
		DbNodeId:   d.NodeID,
		DbPort:     d.Port,
		WebPort:    s.Port,
	}
	c.JSON(http.StatusOK, Response(conf, "success", 200))
}

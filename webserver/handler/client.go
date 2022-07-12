package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientConfig struct {
	Title, Desc string
}

func GetClientConfig(c *gin.Context) {
	conf := ClientConfig{Title: "LemoCoder", Desc: "代码自动生成工具"}
	c.JSON(http.StatusOK, Response(conf, "success", 200))
}

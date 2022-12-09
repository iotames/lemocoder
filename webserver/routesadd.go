package webserver

import (
	"lemocoder/webserver/handler"

	"github.com/gin-gonic/gin"
)

func addRoutes(g *gin.RouterGroup) {
	// Code generated Begin; DO NOT EDIT.
	g.POST("/demo/post", handler.DemoPost)
	// Code generated End; DO NOT EDIT.
}

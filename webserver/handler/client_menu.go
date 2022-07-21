package handler

import (
	"github.com/gin-gonic/gin"
)

type ClientMenuItem struct {
	Layout    bool   `json:"layout"`
	Component string `json:"component"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Redirect  string `json:"redirect"`
	// Icon      string `json:"icon"`
}

func GetClientMenu(c *gin.Context) {
	items := []ClientMenuItem{
		{Path: "/welcome", Name: "首页Wel", Component: "./Welcome", Layout: true},
		{Path: "/excelspider", Name: "Excel爬虫", Component: "./ExcelSpider"},
		{Path: "/", Redirect: "/"},
	}
	c.JSON(200, ResponseItems(items))
}

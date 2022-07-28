package handler

import (
	"lemocoder/database"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	user := getUserModel(c)
	userInfo := map[string]interface{}{
		"Id":      user.ID,
		"Account": user.Account,
		"Name":    user.Name,
		"Avatar":  database.GetDefaultAvatar(),
		"Email":   user.Email,
		"Mobile":  user.Mobile,
	}
	c.JSON(200, Response(userInfo, "success", 200))
}

func Logout(g *gin.Context) {
	user := getUserModel(g)
	if user.ID == 0 {
		ErrorUserNotFound(g)
		return
	}
	user.ResetSalt()
	database.UpdateModel(&user, nil)
	g.JSON(200, ResponseOk("success"))
}

type ClientMenuItem struct {
	// Layout    bool   `json:"layout"`
	// Component string `json:"component"`
	Name string `json:"name"`
	Path string `json:"path"`
	// Redirect string `json:"redirect"`
	// Icon      string `json:"icon"`
}

func GetClientMenu(c *gin.Context) {
	items := []ClientMenuItem{
		{Path: "/welcome", Name: "首页"},          // , Component: "./Welcome"
		{Path: "/tabledemo", Name: "数据表格示例"},    // , Component: "./TableDemo"
		{Path: "/codemaker", Name: "代码生成器"},     // , Component: "./CodeMaker"
		{Path: "/excelspider", Name: "Excel爬虫"}, // , Component: "./ExcelSpider"
	}
	c.JSON(200, ResponseItems(items))
}

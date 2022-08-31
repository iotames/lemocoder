package handler

import (
	"lemocoder/database"
	"lemocoder/initial"

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

func GetClientMenu(c *gin.Context) {
	items := []initial.ClientMenuItem{
		{Path: "/welcome", Name: "首页"},
		{Path: "/tabledemo", Name: "数据表格示例"},
		{Path: "/codemaker", Name: "代码生成器"},
		{Path: "/excelspider", Name: "Excel爬虫"},
		{Path: "/test", Name: "测试"},
	}
	c.JSON(200, ResponseItems(items))
}

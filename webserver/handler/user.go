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
	pages := make([]database.WebPage, 0)
	database.GetAll(&pages, 1000, 1, "project_id = ?", 0)
	var menu []initial.ClientMenuItem
	for _, pg := range pages {
		menu = append(menu, initial.ClientMenuItem{Path: pg.Path, Name: pg.Name})
	}
	c.JSON(200, ResponseItems(initial.GetClientMenu(menu...)))
}

package handler

import (
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userInfo := map[string]interface{}{
		"id":      123568736363656353,
		"account": "QingCoder",
		"name":    "QingCoder",
		"avatar":  "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"email":   "qing@qingcoder.com",
	}
	c.JSON(200, Response(userInfo, "success", 200))
}

func Logout(g *gin.Context) {
	g.JSON(200, ResponseOk("success"))
}

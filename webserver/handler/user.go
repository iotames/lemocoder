package handler

import (
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	result := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"name":        "Serati Ma",
			"avatar":      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
			"userid":      "00000001",
			"email":       "antdesign@alipay.com",
			"signature":   "海纳百川，有容乃大",
			"title":       "交互专家",
			"group":       "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
			"notifyCount": 12,
			"unreadCount": 11,
			"country":     "China",
			"address":     "西湖区工专路 77 号",
			"phone":       "0752-268888888",
		},
	}
	c.JSON(200, result)
}

func Logout(g *gin.Context) {
	g.JSON(200, ResponseOk("success"))
}

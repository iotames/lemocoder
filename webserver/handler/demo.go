package handler

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func DemoPost(c *gin.Context) {
	data := new(map[string]interface{})
	b := c.Bind(data)
	if b != nil {
		c.JSON(200, ResponseFail("请求参数解析错误", 404))
		return
	}
	log.Printf("\n--postDemo: %+v---\n", data)
	time.Sleep(1 * time.Second)
	c.JSON(200, ResponseOk("操作成功")) // Response(data, "success", 200)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonObject map[string]interface{}

func Response(data interface{}, msg string, code int) interface{} {
	return struct {
		Code int
		Msg  string
		Data interface{}
	}{Msg: msg, Code: code, Data: data}
}

func ResponseOk(msg string) interface{} {
	return struct {
		Code int
		Msg  string
		Data JsonObject
	}{Msg: msg, Code: 200, Data: JsonObject{}}
}

func ResponseFail(msg string, code int) interface{} {
	return struct {
		Code int
		Msg  string
		Data JsonObject
	}{Msg: msg, Code: code, Data: JsonObject{}}
}

func ResponseItems(items interface{}) interface{} {
	return struct {
		Code int
		Msg  string
		Data JsonObject
	}{Msg: "success", Code: 200, Data: map[string]interface{}{
		"Items": items,
	}}
}

func ErrorNoPermission(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseFail("NoPermission.您没有权限访问此页面", 400))
}

func ErrorUserNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseFail("user not found", 500))
}
func ErrorNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseFail("无法找到请求对象", http.StatusNotFound))
}
func ErrorArgs(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseFail("请求参数错误", 400))
}
func ErrorServer(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ResponseFail(fmt.Sprintf("服务器系统错误(%v)", err), 500))
}

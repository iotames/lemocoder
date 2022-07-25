package handler

import (
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
	c.JSON(200, ResponseFail("NoPermission.您没有权限访问此页面", 400))
}

func ErrorUserNotFound(c *gin.Context) {
	c.JSON(200, ResponseFail("user not found", 500))
}

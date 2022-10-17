package handler

import (
	"encoding/json"
	"fmt"
	"lemocoder/database"
	"net/http"
	"strings"

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
	switch items.(type) {
	case string:
		return fmt.Sprintf(`{"Code":200,"Msg":"success","Data":{"Items":%s}}`, items)
	default:
		return struct {
			Code int
			Msg  string
			Data JsonObject
		}{Msg: "success", Code: 200, Data: JsonObject{
			"Items": items,
		}}
	}
}

func ItemsIDtoString[T database.IDitem](items []T) (string, error) {
	b, err := json.Marshal(items)
	if err != nil {
		return "", err
	}

	befStr := string(b)
	for _, item := range items {
		fmt.Printf("---item---%+v--", item)
		oldstr := fmt.Sprintf(`"ID":%d`, item.GetID())
		newstr := fmt.Sprintf(`"ID":"%d"`, item.GetID())
		fmt.Println(oldstr, newstr)
		befStr = strings.Replace(befStr, oldstr, newstr, 1)
	}
	return befStr, err
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
func ErrorArgs(c *gin.Context, errs ...error) {
	msg := "请求参数错误"
	if len(errs) > 0 {
		msg += "," + errs[0].Error()
	}
	c.JSON(http.StatusOK, ResponseFail(msg, 400))
}
func ErrorServer(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ResponseFail(fmt.Sprintf("服务器系统错误(%v)", err), 500))
}

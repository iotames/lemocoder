package handler

import (
	"bytes"
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

// ItemsIDtoString 把列表数据的ID字段转为字符串类型，否则JS不支持长整型。会出错。
func ItemsIDtoString[T database.IDitem](items []T) (result string, err error) {
	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(items)
	result = buf.String()
	// fmt.Printf("\n-------%s------\n", result)
	for _, item := range items {
		// fmt.Printf("---item---%+v--", item)
		oldstr := fmt.Sprintf(`"ID":%d`, item.GetID())
		newstr := fmt.Sprintf(`"ID":"%d"`, item.GetID())
		// json数值类型只有 number, 对长整型会丢失精度
		result = strings.Replace(result, oldstr, newstr, 1)
	}
	return
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
	for i, err := range errs {
		if i == 0 {
			msg += "("
		}
		msg += err.Error()
		if i == len(errs)-1 {
			msg += ")"
		} else {
			msg += ","
		}
	}
	// if len(errs) > 0 {
	// 	msg += "," + errs[0].Error()
	// }
	c.JSON(http.StatusOK, ResponseFail(msg, 400))
}
func ErrorServer(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ResponseFail(fmt.Sprintf("服务器系统错误(%v)", err), 500))
}

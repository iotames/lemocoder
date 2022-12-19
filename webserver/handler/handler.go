package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostData map[string]interface{}

func (p *PostData) ParseBody(body io.ReadCloser) error {
	reqBodys, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(reqBodys, p)
}

func (p PostData) ParseTo(result interface{}) error {
	p["ID"] = p.GetID()
	bts, err := json.Marshal(p)
	if err != nil {
		return err
	}
	// "StructSchema":{ ... } json: cannot unmarshal object into Go struct field DataTable.StructSchema of type string
	return util.JsonDecodeUseNumber(bts, result)
}

func (p PostData) GetID() int64 {
	id, ok := p["ID"]
	if !ok {
		return 0
	}
	var result int64
	switch id := id.(type) {
	case float64:
		result = int64(id)
	case string:
		result, _ = strconv.ParseInt(id, 10, 64)
	default:
		result = 0
	}
	return result
}

func (p PostData) Debug() {
	for k, v := range p {
		fmt.Printf("\n----key(%s)--type(%T)---value(%+v)----\n", k, v, v)
	}
}

func getBaseUrl() string {
	w := config.GetWebServer()
	return w.GetAddr()
}

func getUploadsUrl() string {
	return getBaseUrl() + "/uploads"
}

func getUserModel(c *gin.Context) database.User {
	u, exist := c.Get("user")
	if !exist {
		return database.User{}
	}
	return u.(database.User)
}

func CheckBindArgs[T any](args T, c *gin.Context) error {
	err := c.ShouldBind(args) // c.Bind(args)
	if err != nil {
		ErrorArgs(c, err)
	}
	return err
}

func ParsePostData(c *gin.Context) (postData PostData, err error) {
	postData = PostData{}
	err = postData.ParseBody(c.Request.Body)
	if err != nil {
		ErrorServer(c, fmt.Errorf("ParsePostData error:%w", err))
		return
	}
	return
}

func mustFind(c *gin.Context, m database.IModel) error {
	has, err := database.GetModel(m)
	if err != nil {
		ErrorServer(c, err)
		return err
	}
	if !has {
		ErrorNotFound(c)
		err = fmt.Errorf("无法找到请求对象")
	}
	return err
}

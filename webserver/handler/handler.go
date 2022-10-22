package handler

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostData map[string]interface{}

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

func CheckArgs[T any](args T, c *gin.Context) error {
	err := c.Bind(args)
	if err != nil {
		ErrorArgs(c, err)
	}
	return err
}

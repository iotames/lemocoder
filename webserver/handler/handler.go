package handler

import (
	"lemocoder/config"
	"lemocoder/database"

	"github.com/gin-gonic/gin"
)

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

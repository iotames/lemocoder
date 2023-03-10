package handler

import (
	"lemocoder/database"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	AutoLogin bool
	LoginWay  string `binding:"required"` // account, mobile
	Username  string
	Password  string
	Mobile    string
	Captcha   string
}

func Login(c *gin.Context) {
	loginForm := &LoginForm{}
	b := c.Bind(loginForm)
	if b == nil {
		log.Printf("\n--LoginForm: %+v---\n", loginForm)
	} else {
		c.JSON(200, ResponseFail("请求参数解析错误", 404))
		return
	}

	user := new(database.User)
	if loginForm.Username != "" && loginForm.LoginWay == "account" {
		user.Account = loginForm.Username
		database.GetModel(user)
	}
	if loginForm.Mobile != "" && loginForm.LoginWay == "mobile" {
		user.Mobile = loginForm.Mobile
		database.GetModel(user)
	}

	if user.ID == 0 {
		c.JSON(200, ResponseFail("Not found User", 400))
	} else {
		if !user.CheckPassword(loginForm.Password) {
			c.JSON(200, ResponseFail("密码不正确", 400))
			return
		}
		setLoginResponse(c, user)
	}
}

func setLoginResponse(c *gin.Context, u *database.User) {
	jwtInfo := u.GetJwtInfo()
	dt := map[string]interface{}{
		// "department": u.GetDepartment().Name, //resp.Data.Department,
		// "position":   u.GetPosition().Name,
		"Account":   u.Account,
		"Avatar":    u.Avatar,
		"Token":     jwtInfo.Token,
		"ExpiresIn": jwtInfo.Expiresin,
	}
	c.JSON(200, Response(dt, "success", 200))
}

package handler

import (
	"github.com/gin-gonic/gin"
)

type ClientMenuItem struct {
	Layout    bool   `json:"layout"`
	Component string `json:"component"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Redirect  string `json:"redirect"`
	// Icon      string `json:"icon"`
}

func GetMenu(c *gin.Context) {
	items := []ClientMenuItem{
		{Path: "/welcome", Name: "首页", Component: "./Welcome", Layout: true},
		{Path: "/", Redirect: "/"},
	}
	c.JSON(200, ResponseItems(items))
}

// export default [
//     {
//         layout: false,
//         path: '/public/login',
//         // name: 'ligin',
//         component: './public/Login',
//     },
//     {
//         path: '/welcome',
//         name: 'welcome',
//         icon: 'smile',
//         component: './Welcome',
//     },
//     { path: '/', redirect: '/welcome' },
//     { component: './404' },
// ];

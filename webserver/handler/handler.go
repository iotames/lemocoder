package handler

import (
	"lemocoder/config"
)

func getBaseUrl() string {
	w := config.GetWebServer()
	return w.GetAddr()
}

func getUploadsUrl() string {
	return getBaseUrl() + "/uploads"
}

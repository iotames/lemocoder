package handler

import (
	"os"
)

func getBaseUrl() string {
	return os.Getenv("WEB_SERVER_URL")
}

func getUploadsUrl() string {
	return getBaseUrl() + "/uploads"
}

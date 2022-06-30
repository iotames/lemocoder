package main

import (
	"lemocoder/webserver"

	"github.com/joho/godotenv"
)

func main() {
	webserver.New().ListenAndServe()
}

func init() {
	godotenv.Load(".env", "env.default")
}

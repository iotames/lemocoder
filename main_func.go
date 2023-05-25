package main

import (
	"fmt"
	"io/ioutil"
	"lemocoder/util"
	"net/http"
	"os"
)

func stopApp() {
	portStr := os.Getenv("WEB_SERVER_PORT")
	postUrl := fmt.Sprintf("http://127.0.0.1:%s/api/local/stop", portStr)
	// postUrl := "https://httpbin.org/post"
	fmt.Println("----------post url:", postUrl)
	req, err := http.NewRequest("POST", postUrl, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("---Response(%s)----", string(b))
}

func startBrowser() error {
	return util.StartBrowserByUrl("http://127.0.0.1:" + os.Getenv("WEB_SERVER_PORT") + "/client")
}

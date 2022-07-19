package handler

import (
	"lemocoder/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadFileInfo struct {
	Url, Path, ID string
}
type ExcelSpiderForm struct {
	Spider, Title, Uploadfile string
	SheetName                 string `json:"sheet_name"`
	UrlTitle                  string `json:"url_title"`
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("files")
	fileSplit := strings.Split(file.Filename, `.`)
	fileExt := fileSplit[len(fileSplit)-1]
	filename := util.GetRandString(32)
	dst := "runtime/uploads/" + filename + "." + fileExt
	c.SaveUploadedFile(file, dst)
	upf := UploadFileInfo{ID: filename}
	c.JSON(200, Response(upf, "success", 200))
}
func ExcelSpider(c *gin.Context) {
	form := new(ExcelSpiderForm)
	b := c.Bind(form)
	if b != nil {
		c.JSON(200, ResponseFail("请求参数解析错误", 404))
		return
	}
	lockFile := form.Uploadfile + ".lock"
	if util.IsPathExists(lockFile) {
		c.JSON(http.StatusOK, ResponseFail("请耐心等待", 400))
		return
	}
	// f, _ := os.OpenFile(lockFile, os.O_CREATE, 0644)
	// f.Close()
	// time.Sleep(10 * time.Second)
	// err := os.Remove(lockFile)
	// if err != nil {
	// 	log.Println("error:", err)
	// }
	log.Println(form)
	c.JSON(http.StatusOK, ResponseOk("提交成功"))
}

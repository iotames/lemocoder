package handler

import (
	"lemocoder/config"
	"lemocoder/database"
	gen "lemocoder/generator"
	"lemocoder/util"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ClientConfig struct {
	IsLocked                                                                                   bool
	Title, Logo, DbDriver, DbHost, DbName, DbPassword, DbUsername, LoginAccount, LoginPassword string
	DbNodeId, DbPort, WebServerPort                                                            int
}

func GetClientConfig(c *gin.Context) {
	d := config.GetDatabase()
	s := config.GetWebServer()
	a := config.GetApp()
	conf := ClientConfig{
		IsLocked:      util.IsPathExists("app.lock"),
		Title:         a.Title,
		Logo:          a.Logo,
		DbDriver:      d.Driver,
		DbHost:        d.Host,
		DbName:        d.Name,
		DbUsername:    d.Username,
		DbPassword:    d.Password,
		DbNodeId:      d.NodeID,
		DbPort:        d.Port,
		WebServerPort: s.Port,
	}
	c.JSON(http.StatusOK, Response(conf, "success", 200))
}

func ClientInit(c *gin.Context) {
	var err error
	conf := new(ClientConfig)
	c.Bind(conf)
	setDefaultInit(conf)
	if len(conf.LoginPassword) < 6 {
		c.JSON(http.StatusOK, ResponseFail("密码长度过短", 400))
		return
	}
	lockFile := "app.lock"
	if util.IsPathExists(lockFile) {
		c.JSON(http.StatusOK, ResponseFail("请不要重复初始化", 400))
		return
	}
	err = createInitFile(*conf)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("配置文件初始化失败: "+err.Error(), 500))
		return
	}
	err = databaseInit(*conf)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("用户注册失败: "+err.Error(), 500))
		return
	}
	f, _ := os.Create(lockFile)
	f.Close()
	config.LoadEnv()
	c.JSON(http.StatusOK, ResponseOk("success"))
}

func setDefaultInit(conf *ClientConfig) {
	if conf.Logo == "" {
		a := config.GetApp()
		conf.Logo = a.Logo
	}
	if conf.WebServerPort == 0 {
		conf.WebServerPort = config.DEFAULT_WEB_SERVER_PORT
	}
}

func createInitFile(conf ClientConfig) error {
	f, err := os.OpenFile(config.EnvFilepath, os.O_RDWR, 0644)
	if err != nil {
		// open .env File Error
		return err
	}
	err = gen.SetContentByTplFile(config.TplDirPath+"/env.tpl", f, conf)
	if err != nil {
		// create .env File Error
		return err
	}
	f.Close()
	return nil
}

func databaseInit(conf ClientConfig) error {
	database.SetEngine(config.Database{
		Driver:   conf.DbDriver,
		Host:     conf.DbHost,
		Username: conf.DbUsername,
		Password: conf.DbPassword,
		Name:     conf.DbName,
		Port:     conf.DbPort,
		NodeID:   conf.DbNodeId,
	})

	database.CreateTables()
	var err error
	u := database.User{Account: conf.LoginAccount}
	u, err = u.Register(conf.LoginPassword)
	return err
}

func CreateCode(c *gin.Context) {
	fields := []gen.FormFieldSchema{
		{Group: []gen.FormFieldSchema{
			{Component: "ProFormSelect", Name: "useMode", Label: "生效方式"},
			{Component: "ProFormDateRangePicker", Name: "contractTime", Label: "有效期"},
		}},
		{Name: "name", Label: "客户名称", Component: "ProFormText"},
		{Name: "company", Label: "我方公司名称", Component: "ProFormText"},
	}
	createForm := gen.ModalFormSchema{
		Key:    "create",
		Button: gen.ButtonSchema{Type: "primary", Title: "创建"},
		Form:   gen.FormSchema{Title: "添加数据", SubmitUrl: "/api/demo/post", FormFields: fields},
	}
	t := gen.TableSchema{
		ItemDataTypeName: "TestTableItem",
		ItemsDataUrl:     "/api/table/demodata",
		ItemUpdateUrl:    "/api/demo/post",
		ItemCreateUrl:    "/api/demo/post",
		ItemDeleteUrl:    "/api/demo/post",
		Items: []gen.TableItemSchema{
			{DataName: "id", Title: "ID", ColSize: 0.7, Copyable: true, DataType: "number"},
			{DataName: "title", Title: "标题", ColSize: 1, Editable: true, Copyable: true, DataType: "string", Search: true},
			{DataName: "created_at", Title: "创建时间", ValueType: "dateTime", Sorter: true, DataType: "string"},
		},
		ModalForms: []gen.ModalFormSchema{createForm},
	}
	err := t.Create()
	if err != nil {
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}
	c.JSON(200, ResponseOk("SUCCESS"))
}

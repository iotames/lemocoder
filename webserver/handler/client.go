package handler

import (
	"lemocoder/config"
	"lemocoder/database"
	gen "lemocoder/generator"
	"lemocoder/model"
	"lemocoder/util"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetClientConfig(c *gin.Context) {
	conf := config.ClientConfig{}
	(&conf).Load()
	if conf.IsLocked {
		// 应用安装完成后，不可再暴露敏感信息
		conf.DbDriver, conf.DbHost, conf.DbName, conf.DbPassword, conf.DbUsername, conf.LoginAccount, conf.LoginPassword = "", "", "", "", "", "", ""
		conf.DbNodeId, conf.DbPort = 0, 0
	}
	c.JSON(http.StatusOK, Response(conf, "success", 200))
}

func ClientInit(c *gin.Context) {
	var err error
	conf := new(config.ClientConfig)
	c.Bind(conf)
	conf.SetDefaultIfEmpty()
	if len(conf.LoginPassword) < 6 {
		c.JSON(http.StatusOK, ResponseFail("密码长度过短", 400))
		return
	}
	lockFile := "app.lock"
	if util.IsPathExists(lockFile) {
		c.JSON(http.StatusOK, ResponseFail("请不要重复初始化", 400))
		return
	}
	err = conf.Save()
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

func databaseInit(conf config.ClientConfig) error {
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
	fields := []model.FormFieldSchema{
		{Group: []model.FormFieldSchema{
			{Component: "ProFormSelect", Name: "useMode", Label: "生效方式"},
			{Component: "ProFormDateRangePicker", Name: "contractTime", Label: "有效期"},
		}},
		{Name: "name", Label: "客户名称", Component: "ProFormText"},
		{Name: "company", Label: "我方公司名称", Component: "ProFormText"},
	}
	rowFormFields := []model.FormFieldSchema{
		{Name: "id", Label: "ID", Component: "ProFormText"},
		{Name: "title", Label: "Title", Component: "ProFormText"},
	}
	createForm := model.ModalFormSchema{
		Key:    "create",
		Button: model.ButtonSchema{Type: "primary", Title: "创建"},
		Form:   model.FormSchema{Title: "添加数据", SubmitUrl: "/api/demo/post", FormFields: fields},
	}
	editForm := model.ModalFormSchema{
		Key:  "editform1",
		Form: model.FormSchema{Title: "编辑数据", SubmitUrl: "/api/demo/post", FormFields: rowFormFields},
	}
	t := model.TableSchema{
		ItemDataTypeName: "TestTableItem",
		ItemsDataUrl:     "/api/table/demodata",
		ItemUpdateUrl:    "/api/demo/post",
		ItemCreateUrl:    "/api/demo/post",
		ItemDeleteUrl:    "/api/demo/post",
		RowKey:           "id",
		Items: []model.TableItemSchema{
			{DataName: "id", Title: "ID", ColSize: 0.7, Copyable: true, DataType: "int"},
			{DataName: "title", Title: "标题", ColSize: 1, Editable: true, Copyable: true, DataType: "string", Search: true},
			{DataName: "price", Title: "价格", ColSize: 1, Editable: true, Copyable: true, DataType: "float", Search: true},
			{DataName: "inventory", Title: "库存", ColSize: 1, Editable: true, Copyable: true, DataType: "int", Search: true},
			{DataName: "created_at", Title: "创建时间", ValueType: "dateTime", Sorter: true, DataType: "string"},
		},
		ItemOptions: []model.TableItemOptionSchema{
			{Key: "edit", Title: "行编辑", Type: "edit"},
			{Key: "editform1", Title: "表单编辑", Type: "form"},
			{Key: "post1", Title: "标记", Type: "action", Url: "/api/demo/post"},
			{Key: "ret", Title: "跳转", Type: "redirect", Url: "/welcome"},
		},
		ItemForms:       []model.ModalFormSchema{editForm},
		ToolBarForms:    []model.ModalFormSchema{createForm},
		BatchOptButtons: []model.BatchOptButtonSchema{{Title: "批量操作A", Url: "/api/demo/post"}, {Title: "批量操作B", Url: "/api/demo/post"}},
	}
	webpage := database.WebPage{Component: "Test", Path: "/test", Name: "测试菜单1"}
	err := gen.CreateTableClient(t, webpage)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for CreateTableClient:", err)
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}
	err = gen.CreateTableServer(t)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for CreateTableServer:", err)
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}
	c.JSON(200, ResponseOk("SUCCESS"))
}

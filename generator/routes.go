package generator

import (
	"bytes"
	"io"
	"lemocoder/config"
	"lemocoder/model"
	"os"
)

func AddApiRoutes(apiRoutes []model.ApiRoute) error {
	// 读取原文件内容
	f, err := os.OpenFile(config.ServerApiRoutesPath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	var before []byte
	before, err = io.ReadAll(f)
	if err != nil {
		return err
	}
	f.Close()

	// 获取需要新增的内容
	data := map[string]interface{}{
		"Routes": apiRoutes,
	}
	var bf bytes.Buffer
	tplText := `<%{range .Routes}%>
	g.<%{.Method}%>("<%{.Path}%>", <%{.FuncName}%>)
	<%{end}%>// ADD ROUTES`
	err = SetContentByTplText(tplText, data, &bf)
	if err != nil {
		return err
	}

	// 写入变更后的内容
	f, err = os.OpenFile(config.ServerApiRoutesPath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	f.Write(bytes.Replace(before, []byte(`	// ADD ROUTES`), bf.Bytes(), 1))
	return f.Close()
}

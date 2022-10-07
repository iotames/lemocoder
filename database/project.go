package database

// "lemocoder/generator"
// 	"encoding/json"
// 	"fmt"
// 	"lemocoder/util"
// 	"log"
// 	"time"

// https://xorm.io/zh/docs/chapter-02/4.columns/  comment	设置字段的注释（当前仅支持mysql）
type Project struct {
	BaseModel                 `xorm:"extends"`
	Name, Title, Remark, Desc string `xorm:"varchar(64) notnull"`
	ClientFrame, ServerFrame  int    `xorm:"notnull default(0)"`
}

func (d Project) TableName() string {
	return "projects"
}

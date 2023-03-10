package database

// "lemocoder/generator"
// 	"encoding/json"
// 	"fmt"
// 	"lemocoder/util"
// 	"log"
// 	"time"

// https://xorm.io/zh/docs/chapter-02/4.columns/  comment	设置字段的注释（当前仅支持mysql）
type WebPage struct {
	BaseModel                            `xorm:"extends"`
	ProjectID                            int64  `xorm:"notnull default(0) 'project_id'"`
	PageType                             int    `xorm:"notnull default(0)"`
	Name, Title, Remark, Path, Component string `xorm:"varchar(64) notnull"`
	State                                int    `xorm:"SMALLINT notnull default(0)"`
}

const PAGE_STATE_EMPTY = 0
const PAGE_STATE_READY = 1
const PAGE_STATE_CREATED = 2
const PAGE_STATE_BUILT = 3

func (d WebPage) TableName() string {
	return "web_pages"
}

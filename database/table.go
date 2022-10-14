package database

import (
	"lemocoder/generator"
	// 	"encoding/json"
	// 	"fmt"
	// 	"lemocoder/util"
	// 	"log"
	// 	"time"
)

// https://xorm.io/zh/docs/chapter-02/4.columns/  comment	设置字段的注释（当前仅支持mysql）
type DataTable struct {
	BaseModel           `xorm:"extends"`
	PageID              int64                 `xorm:"notnull default(0) 'page_id'"`
	Name, Title, Remark string                `xorm:"varchar(64) notnull"`
	StructSchema        generator.TableSchema `xorm:"TEXT notnull"`
}

// func (d DataTable) GetStructSchema() generator.TableSchema {
// 	ts := generator.TableSchema{}
// 	json.Unmarshal([]byte(d.StructSchema), &ts)
// 	return ts
// }

func (d DataTable) TableName() string {
	return "data_tables"
}

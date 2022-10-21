package database

import (
	"encoding/json"
	"lemocoder/model"
	// 	"fmt"
	// 	"lemocoder/util"
	// 	"log"
	// 	"time"
)

// https://xorm.io/zh/docs/chapter-02/4.columns/  comment	设置字段的注释（当前仅支持mysql）
type DataTable struct {
	BaseModel           `xorm:"extends"`
	PageID              int64  `xorm:"notnull default(0) 'page_id'"`
	Name, Title, Remark string `xorm:"varchar(64) notnull"`
	StructSchema        string `xorm:"TEXT notnull"`
}

func (d DataTable) GetStructSchema() (model.TableSchema, error) {
	var err error
	ts := model.TableSchema{}
	if d.StructSchema != "" {
		err = json.Unmarshal([]byte(d.StructSchema), &ts)
	}
	return ts, err
}

func (d *DataTable) SetStructSchema(ts model.TableSchema) error {
	b, err := json.Marshal(ts)
	d.StructSchema = string(b)
	return err
}

func (d DataTable) TableName() string {
	return "data_tables"
}

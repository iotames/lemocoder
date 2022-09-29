package database

import (
	"lemocoder/generator"
	// 	"encoding/json"
	// 	"fmt"
	// 	"lemocoder/util"
	// 	"log"
	// 	"time"
)

type DataTable struct {
	BaseModel `xorm:"extends"`
	// https://xorm.io/zh/docs/chapter-02/4.columns/  comment	设置字段的注释（当前仅支持mysql）
	Title, Remark, Path, Component string                `xorm:"varchar(64) notnull"`
	StructSchema                   generator.TableSchema `xorm:"TEXT notnull"`
	Name                           string                `xorm:"varchar(32) notnull"`
}

func (d DataTable) TableName() string {
	return "data_tables"
}

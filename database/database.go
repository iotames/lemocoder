package database

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	once   sync.Once
	engine *xorm.Engine
	snode  *snowflake.Node
)

func getNodeId() int64 {
	nidstr := os.Getenv("DB_NODE_ID")
	nodeId, _ := strconv.ParseInt(nidstr, 10, 64)
	return nodeId
}

func GetEngine() *xorm.Engine {
	once.Do(func() {
		engine = newEngine()
	})
	return engine
}

func newEngine() *xorm.Engine {
	dbDriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	var err error
	if dbDriver == "mysql" {
		engine, err = xorm.NewEngine(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname))
	}
	if dbDriver == "sqlite" {
		engine, err = xorm.NewEngine("sqlite3", "./sqlite3.db")
	}
	if err != nil {
		panic(err)
	}
	engineInit(engine)
	return engine
}

func engineInit(engine *xorm.Engine) {
	ormMap := names.GonicMapper{}
	engine.SetMapper(ormMap)
	engine.SetTableMapper(ormMap)
	engine.SetColumnMapper(ormMap)
	engine.ShowSQL(true)
}

func getSnowflakeNode() *snowflake.Node {
	once.Do(func() {
		node, err := snowflake.NewNode(getNodeId())
		if err != nil {
			fmt.Println(err)
			snode = nil
		}
		snode = node
	})
	return snode
}

type BaseModel struct {
	// TODO 分布式ID 雪花算法 https://www.itqiankun.com/article/1565747019
	ID        int64     `xorm:"pk unique"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (b *BaseModel) GenerateID() int64 {
	if b.ID == 0 {
		id := getSnowflakeNode().Generate().Int64()
		if id == 0 {
			panic("Error: getSnowflakeNode().Generate().Int64() == 0")
		}
		b.ID = id
	}
	return b.ID
}

func (b BaseModel) ParseID() snowflake.ID {
	return snowflake.ParseInt64(b.ID)
}

func CreateTables() {
	GetEngine().CreateTables(new(User))
}

func SyncTables() {
	GetEngine().Sync(new(User))
}

package database

import (
	"fmt"
	"lemocoder/config"
	"log"
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
	d := config.GetDatabase()
	return int64(d.NodeID)
}

func GetEngine() *xorm.Engine {
	once.Do(func() {
		engine = newEngine()
	})
	return engine
}

func newEngine() *xorm.Engine {
	db := config.GetDatabase()
	var err error
	if db.Driver == config.DRIVER_MYSQL {
		engine, err = xorm.NewEngine(db.Driver, db.GetDSN())
	}
	if db.Driver == config.DRIVER_SQLITE3 {
		engine, err = xorm.NewEngine(db.Driver, config.SQLITE_FILENAME)
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
	if snode == nil {
		node, err := snowflake.NewNode(getNodeId())
		if err != nil {
			fmt.Println(err)
			snode = nil
		}
		snode = node
	}
	log.Println("---getSnowflakeNode---", snode)
	return snode
}

type IModel interface {
	GenerateID() int64
	ParseID() snowflake.ID
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

func CreateModel(m IModel) (int64, error) {
	m.GenerateID()
	return GetEngine().Insert(m)
}

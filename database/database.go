package database

import (
	"fmt"
	"lemocoder/config"
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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

func getEngine() *xorm.Engine {
	if engine != nil {
		return engine
	}
	once.Do(func() {
		engine = newEngine(nil)
	})
	return engine
}

func SetEngine(db config.Database) {
	engine = newEngine(&db)
}

func newEngine(db *config.Database) *xorm.Engine {
	log.Println("Init newEngine Begin")
	if db == nil {
		db = config.GetDatabase()
	}
	var err error
	if db.Driver == config.DRIVER_SQLITE3 {
		engine, err = xorm.NewEngine(db.Driver, config.SQLITE_FILENAME)
	} else {
		engine, err = xorm.NewEngine(db.Driver, db.GetDSN())
	}
	if err != nil {
		panic(err)
	}
	engineInit(engine)
	log.Println("Init newEngine End")
	return engine
}

func engineInit(engine *xorm.Engine) {
	log.Println("Init engineInit Begin")
	ormMap := names.GonicMapper{}
	engine.SetMapper(ormMap)
	// engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.SetTableMapper(ormMap)
	engine.SetColumnMapper(ormMap)
	engine.ShowSQL(true)
	log.Println("Init engineInit End")
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
	GetID() int64
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

func (b BaseModel) GetID() int64 {
	return b.ID
}

func CreateTables() {
	err := getEngine().CreateTables(new(User))
	if err != nil {
		panic(fmt.Errorf("xorm Error: CreateTables Fail %v", err))
	}
}

func SyncTables() {
	getEngine().Sync(new(User))
}

func GetModel(m IModel) {
	getEngine().Get(m)
}

func CreateModel(m IModel) (int64, error) {
	m.GenerateID()
	return getEngine().Insert(m)
}

func UpdateModel(m IModel, dt map[string]interface{}) (int64, error) {
	modelID := m.GetID() // m.ParseID().Int64()
	if dt == nil {
		return getEngine().ID(modelID).Update(m)
	}
	return getEngine().Table(m).ID(modelID).Update(dt)
}

package database

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/util"
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
	log.Println("New newEngine Begin")
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
	log.Println("New newEngine End")
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
			logger := util.GetLogger()
			logger.Error("Error for database.getSnowflakeNode:", err)
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

func getAllTables() []interface{} {
	return []interface{}{new(User), new(DataTable), new(Project), new(WebPage)}
}

func CreateTables() {
	err := getEngine().CreateTables(getAllTables()...)
	if err != nil {
		panic(fmt.Errorf("error for database.CreateTables:%v", err))
	}
}

func SyncTables() {
	getEngine().Sync(getAllTables()...)
}

func GetModel(m IModel) (bool, error) {
	b, err := getEngine().Get(m)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for database.GetModel:", err)
	}
	return b, err
}

// GetModelWhere 添加复杂条件. 参数 m IModel 各属性必须为零值，否则查询条件会冲突
// GetModelWhere(new(User), "name = ? AND age = ?", "Tom", 19)
func GetModelWhere(m IModel, query interface{}, args ...interface{}) (bool, error) {
	b, err := getEngine().Where(query, args...).Get(m)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for database.GetModel:", err)
	}
	return b, err
}

// GetAll 获取多条记录
// users := make([]Userinfo, 0)
// GetAll(&users, 50, 3, "age > ? or name = ?", 30, "xlw")
func GetAll(rows interface{}, limit, page int, query interface{}, args ...interface{}) error {
	start := (page - 1) * limit
	err := getEngine().Where(query, args...).Limit(limit, start).Find(rows)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for database.GetAll:", err)
	}
	return err
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

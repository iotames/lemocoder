package config

import (
	"fmt"
	"lemocoder/util"
	"os"
	"strconv"
	"text/template"

	"github.com/joho/godotenv"
)

const ENV_PROD = "prod"
const ENV_DEV = "dev"

const DEFAULT_WEB_SERVER_PORT = 8888
const EnvFilepath = ".env"
const ServerHandlerDir = "webserver/handler"
const ServerApiRoutesPath = "webserver/routesadd.go"
const ServerTablesPath = "database/dbtables.go"
const ClientSrcPagesDir = "webclient/src/pages"
const ClientRoutesPath = "webclient/config/routes.ts"
const ClientFilepath = "resource/client/index.html"
const TplDirPath = "resource/templates"
const DRIVER_SQLITE3 = "sqlite3"
const DRIVER_MYSQL = "mysql"
const DRIVER_POSTGRES = "postgres"

const SQLITE_FILENAME = "sqlite3.db"

func LoadEnv() {
	if !util.IsPathExists(EnvFilepath) {
		f, err := os.Create(EnvFilepath)
		if err != nil {
			panic("Create .env Error: " + err.Error())
		}
		f.Close()
	}
	err := godotenv.Load(EnvFilepath, "env.default")
	if err != nil {
		panic("godotenv Error: " + err.Error())
	}
}

type ClientConfig struct {
	IsLocked                                                                                   bool
	Title, Logo, DbDriver, DbHost, DbName, DbPassword, DbUsername, LoginAccount, LoginPassword string
	DbNodeId, DbPort, WebServerPort                                                            int
}

func (c *ClientConfig) Load() ClientConfig {
	d := GetDatabase()
	s := GetWebServer()
	a := GetApp()
	c.IsLocked = util.IsPathExists("app.lock")
	c.Title = a.Title
	c.Logo = a.Logo
	c.DbDriver = d.Driver
	c.DbHost = d.Host
	c.DbName = d.Name
	c.DbUsername = d.Username
	c.DbPassword = d.Password
	c.DbNodeId = d.NodeID
	c.DbPort = d.Port
	c.WebServerPort = s.Port
	return *c
}

func (c *ClientConfig) SetDefaultIfEmpty() {
	if c.Logo == "" {
		a := GetApp()
		c.Logo = a.Logo
	}
	if c.WebServerPort == 0 {
		c.WebServerPort = DEFAULT_WEB_SERVER_PORT
	}
}

func (c ClientConfig) Save() error {
	f, err := os.OpenFile(EnvFilepath, os.O_RDWR, 0644)
	if err != nil {
		// open .env File Error
		return err
	}
	defer f.Close()
	t, err := template.ParseFiles(TplDirPath + "/env.tpl")
	if err != nil {
		return err
	}
	return t.Execute(f, c)
}

type Database struct {
	Driver, Host, Username, Password, Name string
	Port, NodeID                           int
}

func GetDatabase() *Database {
	dbDriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	nodeIdStr := os.Getenv("DB_NODE_ID")
	dbname := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("Error: Fail To Get DB_PORT," + err.Error())
	}
	nodeID, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		panic("Error: Fail To Get DB_NODE_ID," + err.Error())
	}
	return &Database{Driver: dbDriver, Host: host, Username: username, Password: password, Name: dbname, Port: port, NodeID: nodeID}
}

func (d Database) GetAddr() string {
	return fmt.Sprintf("%s:%d", d.Host, d.Port)
}

func (d Database) GetDSN() string {
	dsnMap := map[string]string{
		DRIVER_MYSQL:    fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Name),
		DRIVER_POSTGRES: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", d.Host, d.Username, d.Password, d.Name, d.Port),
	}
	dsn, ok := dsnMap[d.Driver]
	if !ok {
		dsnLen := len(dsnMap)
		ds := make([]string, dsnLen)
		for k, _ := range dsnMap {
			ds = append(ds, k)
		}
		errMsg := fmt.Sprintf("ENV error: DB_DRIVER only Support: %v", ds)
		panic(errMsg)
	}
	return dsn
}

type WebServer struct {
	Port int
}

func GetWebServer() *WebServer {
	portStr := os.Getenv("WEB_SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("Error: Fail To Get WEB_SERVER_PORT," + err.Error())
	}
	return &WebServer{Port: port}
}

func (s WebServer) GetAddr() string {
	return fmt.Sprintf("http://127.0.0.1:%d", s.Port)
}

type App struct {
	Title, Logo, Env, Version string
}

func GetApp() *App {
	runEnv := os.Getenv("APP_ENV")
	if runEnv == "" {
		runEnv = ENV_DEV
	}
	return &App{
		Title:   os.Getenv("APP_TITLE"),
		Logo:    os.Getenv("APP_LOGO"),
		Env:     runEnv,
		Version: os.Getenv("APP_VERSION"),
	}
}

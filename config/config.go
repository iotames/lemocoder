package config

import (
	"fmt"
	"os"
	"strconv"
)

const DRIVER_SQLITE3 = "sqlite3"
const DRIVER_MYSQL = "mysql"

const SQLITE_FILENAME = "sqlite3.db"

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
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Name)
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

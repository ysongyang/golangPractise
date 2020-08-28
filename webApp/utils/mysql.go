package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//定义sql全局变量
var (
	Db  *sql.DB
	err error
)

//数据库配置
const (
	driverName = "mysql"
	USERNAME   = "root"
	PASSWORD   = "root"
	NETWORK    = "tcp"
	SERVER     = "localhost"
	PORT       = 3306
	DATABASE   = "golang"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	Db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Panicln("Open mysql failed：", err.Error())
		return
	}
}

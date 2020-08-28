package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golangPractise/mallApi/config"
	"log"
)

//定义sql全局变量
var (
	Db  *sql.DB
	err error
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.USERNAME, config.PASSWORD, config.NETWORK, config.SERVER, config.PORT, config.DATABASE)
	Db, err = sql.Open(config.DriverName, dataSourceName)
	if err != nil {
		log.Panicln("Open mysql failed：", err.Error())
		return
	}
	Db.SetMaxOpenConns(0)
	Db.SetMaxIdleConns(0)
}

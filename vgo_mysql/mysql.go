package vgo_mysql

import (
	"fmt"
	"log"

	"github.com/PyreneGitHub/vgo/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
	err := initDb()
	if err != nil {
		log.Println("init db err", err)
		return
	}
}

func initDb() (err error) {
	Client := &conf.Vgo_client{}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.MysqlConfig.UserName, conf.MysqlConfig.Passwd,
		conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.Database)
	database, err := sqlx.Open("mysql", dns)
	if err != nil {
		log.Println("open mysql failed, err:%v ", err)
		return
	}

	Client.Vgo_db = database
	log.Println("connect to mysql succ")
	return
}

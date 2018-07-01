package main

import (
	"log"

	"github.com/PyreneGitHub/vgo/conf"

	"github.com/PyreneGitHub/vgo/read_conf"
)

var config conf.Config

func ConfigAll() error {
	conf, err := read_conf.ReadFile_Conf("D:/project/src/vgo/demo/app/appconf.ini")
	if err != nil {
		log.Println("err", err)
		return err
	}
	config.MysqlConf.UserName = conf.String("mysql.db.UserName")
	config.MysqlConf.Passwd = conf.String("mysql.db.Passwd")
	config.MysqlConf.Host = conf.String("mysql.db.Host")
	config.MysqlConf.Port, err = conf.Int("mysql.db.Port")
	if err != nil {
		log.Println("read mysqlport err", err)
		return err
	}
	config.MysqlConf.Database = conf.String("mysql.db.Database")

	config.EtcdConf.Addr = conf.String("etcd.etcd.Addr")
	config.EtcdConf.Timeout, err = conf.Int("etcd.etcd.Timeout")
	if err != nil {
		log.Println("read etcd timeout", err)
		return err
	}

	config.RedisConf.RedisAddr = conf.String("redis.r.RedisAddr")
	config.RedisConf.RedisMaxActive, err = conf.Int("redis.r.RedisMaxActive")
	if err != nil {
		log.Println("read redis maxactive err", err)
		return err
	}
	config.RedisConf.RedisIdleTimeout, err = conf.Int("redis.r.RedisIdleTimeout")
	if err != nil {
		log.Println("read redis timeout")
		return err
	}
	config.RedisConf.RedisMaxIdle, err = conf.Int("redis.r.RedisMaxIdle")
	if err != nil {
		log.Println("read redis maxidle err", err)
		return err
	}
	return err
}

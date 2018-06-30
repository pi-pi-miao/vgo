package conf

import (
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

func init() {
	_ := RedisConf{
		RedisAddr:        nil,
		RedisMaxActive:   nil,
		RedisIdleTimeout: nil,
		RedisMaxIdle:     nil,
	}
	_ := MysqlConfig{
		UserName: nil,
		Passwd:   nil,
		Host:     nil,
		Port:     nil,
		Database: nil,
	}
	_ := EtcdConf{
		Addr:    nil,
		Timeout: nil,
	}
}

type (
	Config struct {
		MysqlConf  MysqlConfig
		EtcdConf   EtcdConf
		RedisConf  RedisConf
		Vgo_client Vgo_client
	}
	MysqlConfig struct {
		UserName string
		Passwd   string
		Host     string
		Port     int
		Database string
	}

	EtcdConf struct {
		Addr    string
		Timeout int64
	}

	RedisConf struct {
		RedisMaxIdle     int
		RedisMaxActive   int
		RedisIdleTimeout int
		RedisAddr        string
	}
	Vgo_client struct {
		Etcdcli   *etcd_client.Client
		Vgo_db    *sqlx.DB
		RedisPool *redis.Pool
	}
)

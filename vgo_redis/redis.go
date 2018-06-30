package vgo_redis

import(
	"github.com/garyburd/redigo/redis"
	"time"
	"log"
	"vgo/conf"
)

func init(){
	err:=initRedis()
	if err!=nil{
		log.Println("init redis err",err)
		return
	}
}

func initRedis() (err error) {
	Client:=&conf.Vgo_client{}
	Client.RedisPool = &redis.Pool{
		MaxIdle:     conf.RedisConf.RedisMaxIdle,
		MaxActive:   conf.RedisConf.RedisMaxActive,
		IdleTimeout: time.Duration(conf.RedisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.RedisConf.RedisAddr)
		},
	}

	conn := Client.RedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		log.Println("ping redis failed, err:%v", err)
		return
	}

	return
}
package vgo_etcd

import (
	"log"
	"time"

	"github.com/PyreneGitHub/vgo/conf"
	etcd_client "github.com/coreos/etcd/clientv3"
)

func init() {
	err := initEtcd()
	if err != nil {
		log.Println("init etcd err", err)
		return
	}
}

func initEtcd() (err error) {
	Client := &conf.Vgo_client{}
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{conf.EtcdConf.Addr},
		DialTimeout: time.Duration(conf.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		log.Println("connect etcd failed, err:", err)
		return
	}
	Client.Etcdcli = cli
	log.Println("init etcd succ")
	return
}

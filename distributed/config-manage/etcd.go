package config_manage

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/client"
	"log"
	"time"
)

var configPath = `/configs/remote_config.json`
var kapi client.KeysAPI

type ConfigStruct struct {
	Addr           string `json:"addr"`
	AesKey         string `json:"aesKey"`
	HTTPS          bool   `json:"https"`
	Secret         string `json:"secret"`
	PrivateKeyPath string `json:"privateKeyPath"`
	CertFilePath   string `json:"certFilePath"`
}

var appConfig ConfigStruct

func init() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi = client.NewKeysAPI(c)
	initConfig()
}

func watchAndUpdate() {
	w := kapi.Watcher(configPath, nil)
	go func() {
		// watch 该节点下的每次变化
		resp, err := w.Next(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("new value is", resp.Node.Value)

		err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func initConfig() {
	resp, err := kapi.Get(context.Background(), configPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() ConfigStruct {
	return appConfig
}

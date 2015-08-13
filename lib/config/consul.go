package config

import (
	"fmt"
	"news_worker/lib/utils"
	"os"

	consul "github.com/armon/consul-api"
)

var (
	consulEnv = "CONSUL_MASTER_IP"
	kv        *consul.KV
)

// Start app session start for consul
func Start() {
	// Get a new client, with KV endpoints
	client, _ := consul.NewClient(getConsulMasterIP())
	kv = client.KV()

	GetKV("mongodb_cluster1")
}

// GetKV get k/v pair
func GetKV(key string) string {
	// Lookup the pair
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		panic(err)
	}
	utils.Info(fmt.Sprintf("KV: %v", string(pair.Value)))

	return string(pair.Value)
}

// PutValue put/save k/v to consul
func PutValue(key, value string) {
	// PUT a new KV pair
	p := &consul.KVPair{Key: "foo", Value: []byte("test")}
	resp, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	utils.Info(fmt.Sprintf("put success! %v", resp))
}

func getConsulMasterIP() *consul.Config {
	e := os.Getenv(consulEnv)
	if e == "" {
		panic("consul not found!")
	}
	serverIP := fmt.Sprintf("%s:9200", e)
	config := consul.DefaultConfig()
	config.Address = serverIP
	fmt.Println(serverIP)
	return config
}

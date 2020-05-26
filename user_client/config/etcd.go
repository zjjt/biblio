package config

import (
	"os"
	"strconv"
)

type EtcdConfig interface {
	GetPort() int
	GetHost() string
}

// defaultEtcdConfig
type defaultEtcdConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// GetPort Etcd
func (c defaultEtcdConfig) GetPort() int {
	return c.Port
}

// GetHost Etcd
func (c defaultEtcdConfig) GetHost() string {
	return c.Host
}

//GetETCDConfig etcd config
func GetETCDConfig() EtcdConfig {

	port, _ := strconv.Atoi(os.Getenv("ETCD_PORT"))
	host := os.Getenv("ETCD_HOST")

	etcd := defaultEtcdConfig{Host: host, Port: port}
	return etcd
}

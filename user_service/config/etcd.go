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
	//Comment when in production
	os.Setenv("ETCD_PORT", "2379")
	os.Setenv("ETCD_HOST", "127.0.0.1")
	port, _ := strconv.Atoi(os.Getenv("ETCD_PORT"))
	etcd := defaultEtcdConfig{Host: os.Getenv("ETCD_HOST"), Port: port}
	return etcd
}

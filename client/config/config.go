package config

import (
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

type Config struct {
	Addr string
	Port int
}

var configInstance *Config

func GetConfig() *Config {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if configInstance == nil {
			configInstance = &Config{
				Addr: "http://localhost",
				Port: 8080,
			}
		}
	}

	return configInstance
}

func (c *Config) GetURL() string {
	return c.Addr + ":" + strconv.Itoa(c.Port)
}

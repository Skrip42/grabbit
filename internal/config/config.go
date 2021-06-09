package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
    Port      int    `yaml:"port"`
    RabbitDSN string `yaml:"rabbit"`
}

var (
    instance *Config
    once sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        instance = &Config{}
        yamlFile, err := ioutil.ReadFile("./config/config.yaml")
        if err != nil {
            log.Fatal("cannot read config.yaml", err.Error())
        }
        err = yaml.Unmarshal(yamlFile, instance)
        if err != nil {
            log.Fatal("cannot unmarshal config.yaml", err.Error())
        }
    })
    return instance
}

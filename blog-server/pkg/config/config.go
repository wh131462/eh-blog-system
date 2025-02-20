package config

import (
    "github.com/spf13/viper"
    "sync"
)

var (
    once sync.Once
    C    *Config
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Hexo     HexoConfig
}

type ServerConfig struct {
    Port string `yaml:"port"`
    Env  string `yaml:"env"`
}

func Init(path string) {
    once.Do(func() {
        viper.SetConfigFile(path)
        if err := viper.ReadInConfig(); err != nil {
            panic(err)
        }
        viper.Unmarshal(&C)
    })
}
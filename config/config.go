package config

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
	"time"

	"go-gin-frame/pkg/env"
	"go-gin-frame/pkg/file"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	MySQL struct {
		Addr            string        `yaml:"addr"`
		User            string        `yaml:"user"`
		Pass            string        `yaml:"pass"`
		Name            string        `yaml:"name"`
		MaxOpenConn     int           `yaml:"maxOpenConn"`
		MaxIdleConn     int           `yaml:"maxIdleConn"`
		ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
	} `yaml:"mysql"`

	Redis struct {
		Addr         string `yaml:"addr"`
		Pass         string `yaml:"pass"`
		Db           int    `yaml:"db"`
		MaxRetries   int    `yaml:"maxRetries"`
		PoolSize     int    `yaml:"poolSize"`
		MinIdleConns int    `yaml:"minIdleConns"`
	} `yaml:"redis"`

	Mail struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		To   string `yaml:"to"`
	} `yaml:"mail"`

	HashIds struct {
		Secret string `yaml:"secret"`
		Length int    `yaml:"length"`
	} `yaml:"hashids"`

	Language struct {
		Local string `yaml:"local"`
	} `yaml:"language"`
}

var (
	//go:embed dev_config.yaml
	devConfig []byte

	//go:embed fat_config.yaml
	fatConfig []byte

	//go:embed uat_config.yaml
	uatConfig []byte

	//go:embed pro_config.yaml
	proConfig []byte
)

func InitConfig() {
	var r io.Reader

	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfig)
	case "fat":
		r = bytes.NewReader(fatConfig)
	case "uat":
		r = bytes.NewReader(uatConfig)
	case "pro":
		r = bytes.NewReader(proConfig)
	default:
		r = bytes.NewReader(fatConfig)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(env.Active().Value() + "_config")
	viper.AddConfigPath("./config")

	configFile := "./config/" + env.Active().Value() + "_config.yaml"
	_, ok := file.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}

package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	Postgres DBConfig `mapstructure:"postgres"`
}

type DBConfig struct {
	Hostname string `mapstructure:"hostname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"db"`
}

var (
	basepath string
	conf     *ApplicationConfig
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(b)
}

func NewConfig() *ApplicationConfig {
	once := new(sync.Once)
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(basepath)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %s", err))
		}

		if err := viper.Unmarshal(&conf); err != nil {
			panic(fmt.Errorf("vailed to load config %s", err))
		}
	})
	return conf
}

package config

import (
	"bytes"
	_ "embed"
	"errors"

	"github.com/apex/log"
	"github.com/spf13/viper"
)

//go:embed default.yml
var defaultConfig []byte

type Config struct {
	// The address to listen on for HTTP requests.
	Address string
	// The port to listen on for HTTP requests.
	Port int
}

type ConfigReader struct {
	l          log.Interface
	ConfigName string
}

func New(l log.Interface, name string) *ConfigReader {
	viper.AddConfigPath(".")
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	return &ConfigReader{
		l:          l,
		ConfigName: name,
	}
}

func (c *ConfigReader) ParseConfig() (*Config, error) {
	conf := &Config{}
	c.l.Info("Reading config...")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			c.l.Warn("Config does not exist, creating...")
			if err := viper.ReadConfig(bytes.NewReader(defaultConfig)); err != nil {
				return nil, err
			}
			if err := viper.SafeWriteConfig(); err != nil {
				return nil, err
			}
			return nil, errors.New("Default configuration has been created, please edit it")
		} else {
			return nil, err
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

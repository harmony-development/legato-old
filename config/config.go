// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package config

import (
	_ "embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

//go:embed default.yml
var defaultConfig []byte

type Config struct {
	// The address to listen on for HTTP requests.
	Address string
	// The port to listen on for HTTP requests.
	Port           int
	PublicKeyPath  string `yaml:"public-key-path"`
	PrivateKeyPath string `yaml:"private-key-path"`
	Debug          Debug
	Database       Database
	Epheremal      Epheremal
}

type Debug struct {
	RespondWithErrors bool `yaml:"respond-with-errors"`
	LogErrors         bool `yaml:"log-errors"`
}

type Database struct {
	Backend  PersistBackend
	Postgres *PostgresConfig
}

type Epheremal struct {
	Backend EpheremalBackend
	Redis   *RedisConfig
}

type RedisConfig struct {
	Hosts    []string
	Password string
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DB       string
}

type ConfigReader struct {
	ConfigName string
}

func New(name string) *ConfigReader {
	return &ConfigReader{
		ConfigName: name + ".yaml",
	}
}

func (c *ConfigReader) ParseConfig() (*Config, error) {
	conf := &Config{}

	dat, err := ioutil.ReadFile(c.ConfigName)
	if err != nil {
		if os.IsNotExist(err) {
			err := ioutil.WriteFile(c.ConfigName, defaultConfig, 0o660)
			if err != nil {
				return nil, fmt.Errorf("failed to write default config: %+w", err)
			}

			return nil, errors.New("default configuration has been created, please edit it")
		}

		return nil, fmt.Errorf("failed to read config file: %+w", err)
	}
	err = yaml.Unmarshal(dat, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %+w", err)
	}

	return conf, nil
}

func (c *ConfigReader) WatchConfig(onChange func(fsnotify.Event), onError func(error)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				onChange(event)
			case err := <-watcher.Errors:
				onError(err)
			}
		}
	}()

	if err := watcher.Add(c.ConfigName); err != nil {
		return err
	}

	return nil
}

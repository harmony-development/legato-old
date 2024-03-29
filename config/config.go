// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package config

import (

	// for embedding default config.
	_ "embed"
	"errors"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/harmony-development/legato/errwrap"
	"gopkg.in/yaml.v3"
)

// nolint
//go:embed default.yml
var defaultConfig []byte

type Config struct {
	// The address to listen on for HTTP requests.
	Address string
	// The port to listen on for HTTP requests.
	Port           int
	PublicKeyPath  string `yaml:"publicKeyPath"`
	PrivateKeyPath string `yaml:"privateKeyPath"`
	AuthIDLength   int    `yaml:"authIdLength"`
	IDStart        int64  `yaml:"idStart"`
	Debug          Debug
	Database       Database
	Epheremal      Epheremal
}

type Debug struct {
	RespondWithErrors bool `yaml:"respondWithErrors"`
	LogErrors         bool `yaml:"logErrors"`
}

type Database struct {
	Backend  string
	Postgres *PostgresConfig
	SQLite   *SQLiteConfig
}

type Epheremal struct {
	Backend string
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

type SQLiteConfig struct {
	File string
}

type Reader struct {
	ConfigName string
}

var ErrDefaultConfigCreated = errors.New("default configuration has been created, please edit it")

func New(name string) *Reader {
	return &Reader{
		ConfigName: name + ".yaml",
	}
}

func (c *Reader) ParseConfig() (*Config, error) {
	conf := &Config{}

	dat, err := os.ReadFile(c.ConfigName)
	if err != nil {
		if os.IsNotExist(err) {
			// stdlib doesn't have any permission bit enums so a "magic" number is ok here
			// nolint
			err := os.WriteFile(c.ConfigName, defaultConfig, 0o660)
			if err != nil {
				return nil, errwrap.Wrap(err, "failed to write default config")
			}

			return nil, ErrDefaultConfigCreated
		}

		return nil, errwrap.Wrap(err, "failed to read config file")
	}

	if err := yaml.Unmarshal(dat, conf); err != nil {
		return nil, errwrap.Wrap(err, "failed to parse config file")
	}

	return conf, nil
}

func (c *Reader) WatchConfig(onChange func(fsnotify.Event), onError func(error)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errwrap.Wrap(err, "failed to start fsnotify watcher")
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
		return errwrap.Wrap(err, "failed to add config to watcher")
	}

	return nil
}

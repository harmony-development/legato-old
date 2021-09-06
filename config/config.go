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

	"github.com/apex/log"
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
	Database       Database
	Epheremal      Epheremal
}

type DatabaseBackend int

const (
	Postgres DatabaseBackend = iota
)

func (e *DatabaseBackend) UnmarshalText(text []byte) error {
	switch string(text) {
	case "postgres":
		*e = Postgres
		return nil
	default:
		return errors.New("database backend must be one of [postgres]")
	}
}

type EpheremalBackend int

const (
	Redis EpheremalBackend = iota
)

func (e *EpheremalBackend) UnmarshalText(text []byte) error {
	switch string(text) {
	case "redis":
		*e = Redis
		return nil
	default:
		return errors.New("database backend must be one of [postgres]")
	}
}

type Database struct {
	Backend  DatabaseBackend
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
	l          log.Interface
	ConfigName string
}

func New(l log.Interface, name string) *ConfigReader {
	return &ConfigReader{
		l:          l,
		ConfigName: name + ".yaml",
	}
}

func (c *ConfigReader) ParseConfig() (*Config, error) {
	conf := &Config{}
	c.l.Info("Reading config...")

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

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

// TODO: get this to use the actual registered backends list
var (
	persistBackends   = StringSet{}
	ephemeralBackends = StringSet{}
)

func init() {
	persistBackends.Add(
		"postgres",
	)
	ephemeralBackends.Add(
		"bigcache",
		"redis",
	)
}

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

type PersistBackend string

func (e *PersistBackend) UnmarshalText(text []byte) error {
	ok := persistBackends.Has(string(text))
	if !ok {
		return fmt.Errorf("persist backend must be one of: %v", persistBackends.Values())
	}
	*e = PersistBackend(text)
	return nil
}

type EpheremalBackend string

func (e *EpheremalBackend) UnmarshalText(text []byte) error {
	ok := ephemeralBackends.Has(string(text))
	if !ok {
		return fmt.Errorf("ephemeral backend must be one of: %v", ephemeralBackends.Values())
	}
	*e = EpheremalBackend(text)
	return nil
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

type StringSet map[string]struct{}

func (set StringSet) Has(s string) bool {
	_, ok := set[s]
	return ok
}

func (set StringSet) Add(vals ...string) {
	for _, v := range vals {
		set[v] = struct{}{}
	}
}

func (set StringSet) Values() []string {
	ret := []string{}
	for k := range set {
		ret = append(ret, k)
	}
	return ret
}

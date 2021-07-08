package config

import (
	"errors"
	"io"
	"os"

	"github.com/creasty/defaults"
	"github.com/ztrue/tracerr"
	"gopkg.in/yaml.v3"
)

// go:embed default.yml
var defaultConf []byte // nolint:gochecknoglobals

type CassandraConfig struct {
	NumRetries int `default:"3"`
}

type ServerConfig struct {
	Host string `default:"0.0.0.0"`
	Port int    `default:"2289"`

	Cassandra CassandraConfig
}

type Config struct {
	Server ServerConfig
}

func Load(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.WriteFile(path, defaultConf, 0644)
		}

		return nil, tracerr.Wrap(err)
	}

	return LoadConfig(configFile)
}

func LoadConfig(handle io.Reader) (*Config, error) {
	cfg := Config{}
	if err := defaults.Set(&cfg); err != nil {
		return nil, tracerr.Wrap(err)
	}

	content, err := io.ReadAll(handle)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &cfg, nil
}

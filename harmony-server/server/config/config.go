package config

import (
	"github.com/spf13/viper"
)

// Config is the overall configuration for the auth service
type Config struct {
	Server ServerConf
	DB     DBConf
	Sentry SentryConf
}

// ServerConf is the servers configuration
type ServerConf struct {
	Port string
}

// DBConf is the config for the database
type DBConf struct {
	Host     string
	Port     int
	User     string
	Password string
	SSL      bool
}

// SentryConf is the config for sentry
type SentryConf struct {
	Dsn              string
	AttachStacktrace bool
	Enabled          bool
}

// Load reads a config file (JSON format)
func Load() (*Config, error) {
	defaultCFG := Config{
		Server: ServerConf{
			Port: ":2289",
		},
		DB: DBConf{
			Host: "127.0.0.1",
			Port: 5432,
			SSL:  false,
		},
		Sentry: SentryConf{
			Dsn:              "",
			AttachStacktrace: true,
			Enabled:          true,
		},
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetDefault("instanceserver", defaultCFG)
	if err := viper.ReadInConfig(); err != nil {
		if err := viper.SafeWriteConfig(); err != nil {
			return nil, err
		}
		return nil, err
	}
	var cfg Config
	if err := viper.UnmarshalKey("instanceserver", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

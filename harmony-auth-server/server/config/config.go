package config

import (
	"time"

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
	Port               string
	InstanceAPIVersion string
	AvatarPath         string
	UsernameLenMin     int
	UsernameLenMax     int
	PassLenMin         int
	PassLenMax         int
	PassRegex          string
	UserIDLength       int
	SessionLength      int
	SessionExpire      time.Duration
	SessionCacheMax    int
	IDCacheMax         int
	AvatarQuality      int
	AvatarWidth        int
	AvatarHeight       int
	AvatarCrop         bool
	TLS                TLSConf
}

// TLSConf contains the configurations to use for TLS
type TLSConf struct {
	Enabled  bool
	CertPath string
	KeyPath  string
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
			Port:               ":2289",
			InstanceAPIVersion: "v1",
			AvatarPath:         "./avatars",
			UsernameLenMin:     2,
			UsernameLenMax:     48,
			PassLenMin:         5,
			PassLenMax:         128,
			PassRegex:          "",
			UserIDLength:       16,
			SessionLength:      16,
			SessionExpire:      24 * time.Hour,
			SessionCacheMax:    500000,
			IDCacheMax:         500000,
			AvatarQuality:      60,
			AvatarWidth:        128,
			AvatarHeight:       128,
			AvatarCrop:         true,
			TLS: TLSConf{
				Enabled:  false,
				CertPath: "",
				KeyPath:  "",
			},
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
	viper.SetDefault("authserver", defaultCFG)
	if err := viper.ReadInConfig(); err != nil {
		if err := viper.SafeWriteConfig(); err != nil {
			return nil, err
		}
		return nil, err
	}
	var cfg Config
	if err := viper.UnmarshalKey("authserver", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

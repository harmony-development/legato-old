package config

import (
	"errors"
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
	PasswordPolicy     PasswordPolicy
	UsernamePolicy     UsernamePolicy
}

// UsernamePolicy contains the settings for password safety
type UsernamePolicy struct {
	MinLength int
	MaxLength int
}

// PasswordPolicy contains the settings for password safety
type PasswordPolicy struct {
	MinLength    int
	MaxLength    int
	MinCapital   int
	MinLowercase int
	MinNumbers   int
	MinSpecial   int
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
			UsernamePolicy: UsernamePolicy{
				MinLength: 5,
				MaxLength: 32,
			},
			PasswordPolicy: PasswordPolicy{
				MinLength:    5,
				MaxLength:    50,
				MinCapital:   1,
				MinLowercase: 1,
				MinNumbers:   1,
				MinSpecial:   0,
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
		return nil, errors.New("config file not found, default config generated at config.yml. Restart server to see effects")
	}
	var cfg Config
	if err := viper.UnmarshalKey("authserver", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

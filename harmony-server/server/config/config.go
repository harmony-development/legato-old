package config

import (
	"github.com/spf13/viper"
	"github.com/thanhpk/randstr"
)

// Config is the overall configuration for the auth service
type Config struct {
	Server ServerConf
	DB     DBConf
	Sentry SentryConf
}

// ServerConf is the servers configuration
type ServerConf struct {
	SnowflakeStart   int64
	Port             string
	Identity         string
	ImagePath        string
	GuildPicturePath string
	PublicKeyPath    string
	PrivateKeyPath   string
	MaxAttachments   int
	GetMessageCount  int
	OwnerCacheMax    int
	SessionCacheMax  int
	LogErrors        bool
	UsernamePolicy   UsernamePolicy
	PasswordPolicy   PasswordPolicy
}

type UsernamePolicy struct {
	MinLength int
	MaxLength int
}

type PasswordPolicy struct {
	MinLength  int
	MaxLength  int
	MinLower   int
	MinUpper   int
	MinNumbers int
	MinSymbols int
}

// DBConf is the config for the database
type DBConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Models   string
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
			Port:             ":2289",
			Identity:         randstr.Hex(16), // this is what prevents people from using auth tokens on other instances
			ImagePath:        "images",
			GuildPicturePath: "guild-pictures",
			PrivateKeyPath:   "harmony-key.pem",
			PublicKeyPath:    "harmony-key.pub",
			MaxAttachments:   1,
			GetMessageCount:  30,
			OwnerCacheMax:    5096,
			SessionCacheMax:  5096,
			LogErrors:        true,
			SnowflakeStart:   0,
			UsernamePolicy: UsernamePolicy{
				MinLength: 2,
				MaxLength: 32,
			},
			PasswordPolicy: PasswordPolicy{
				MinLength:  5,
				MaxLength:  256,
				MinLower:   1,
				MinUpper:   1,
				MinNumbers: 1,
				MinSymbols: 0,
			},
		},
		DB: DBConf{
			Host:   "127.0.0.1",
			Port:   5432,
			SSL:    false,
			Models: "sql/schemas/models.sql",
		},
		Sentry: SentryConf{
			Dsn:              "",
			AttachStacktrace: true,
			Enabled:          false,
		},
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetDefault("InstanceServer", defaultCFG)
	if err := viper.ReadInConfig(); err != nil {
		if err := viper.SafeWriteConfig(); err != nil {
			return nil, err
		}
		return nil, err
	}
	var cfg Config
	if err := viper.UnmarshalKey("InstanceServer", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

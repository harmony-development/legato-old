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
	SnowflakeStart                    int64
	Port                              int
	FlatfileMediaPath                 string
	PublicKeyPath                     string
	PrivateKeyPath                    string
	MaxAttachments                    int
	GetMessageCount                   int
	OwnerCacheMax                     int
	SessionCacheMax                   int
	SessionDuration                   time.Duration
	LogErrors                         bool
	LogRequests                       bool
	RespondWithErrors                 bool
	NonceLength                       int
	GuildLeaveNotificationQueueLength int
	Avatar                            Avatar
	UsernamePolicy                    UsernamePolicy
	PasswordPolicy                    PasswordPolicy
}

type Avatar struct {
	Width   int
	Height  int
	Quality int
	Crop    bool
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
	DBName   string
	Models   string
	SSL      bool
}

// SentryConf is the config for sentry
type SentryConf struct {
	Dsn              string
	AttachStacktrace bool
	Enabled          bool
}

var DefaultConf = Config{
	Server: ServerConf{
		Port:                              2289,
		FlatfileMediaPath:                 "media",
		PrivateKeyPath:                    "harmony-key.pem",
		PublicKeyPath:                     "harmony-key.pub",
		MaxAttachments:                    1,
		GetMessageCount:                   30,
		OwnerCacheMax:                     5096,
		SessionCacheMax:                   5096,
		SessionDuration:                   48 * time.Hour,
		LogErrors:                         true,
		LogRequests:                       true,
		RespondWithErrors:                 false,
		SnowflakeStart:                    0,
		NonceLength:                       32,
		GuildLeaveNotificationQueueLength: 64,
		Avatar: Avatar{
			Width:   256,
			Height:  256,
			Quality: 50,
			Crop:    true,
		},
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
		DBName: "harmony",
		Models: "sql/schemas/models.sql",
	},
	Sentry: SentryConf{
		Dsn:              "",
		AttachStacktrace: true,
		Enabled:          false,
	},
}

// Load reads a config file (JSON format)
func Load() (*Config, error) {
	defaultCFG := DefaultConf
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
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	log.Println(aurora.Yellow("⌛ Reloading config..."))
	// 	if err := viper.UnmarshalKey("InstanceServer", &cfg); err != nil {
	// 		log.Println("failed to read reloaded config", err)
	// 		return
	// 	}
	// 	log.Println(aurora.Green("✔ Config reloaded successfully").Bold())
	// })
	// viper.WatchConfig()
	return &cfg, nil
}

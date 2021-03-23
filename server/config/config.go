package config

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/creasty/defaults"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/ztrue/tracerr"
)

type AvatarPolicy struct {
	Width   int  `hcl:"Width,optional" default:"256"`
	Height  int  `hcl:"Height,optional" default:"256"`
	Quality int  `hcl:"Quality,optional" default:"50"`
	Crop    bool `hcl:"Crop,optional" default:"true"`
}

type UsernamePolicy struct {
	MinLength int `hcl:"MinLength,optional" default:"2"`
	MaxLength int `hcl:"MaxLength,optional" default:"32"`
}

type PasswordPolicy struct {
	MinLength  int `hcl:"MinLength,optional" default:"5"`
	MaxLength  int `hcl:"MaxLength,optional" default:"256"`
	MinLower   int `hcl:"MinLower,optional" default:"1"`
	MinUpper   int `hcl:"MinUpper,optional" default:"1"`
	MinNumbers int `hcl:"MinNumbers,optional" default:"1"`
	MinSymbols int `hcl:"MinSymbols,optional" default:"0"`
}

type AttachmentPolicy struct {
	MaximumCount int `hcl:"MaximumCount,optional" default:"10"`
}

type DebugPolicy struct {
	LogErrors                  bool `hcl:"LogErrors,optional" default:"true"`
	LogRequests                bool `hcl:"LogRequests,optional" default:"true"`
	RespondWithErrors          bool `hcl:"RespondWithErrors,optional" default:"false"`
	ResponseErrorsIncludeTrace bool `hcl:"ResponseErrorsIncludeTrace" default:"true"`
	VerboseStreamHandling      bool `hcl:"VerboseStreamHandling" default:"false"`
}

type SessionPolicy struct {
	Duration time.Duration `hcl:"Duration,optional" default:"172800000000000"`
}

type CachePolicy struct {
	Owner       int `hcl:"Owner,optional" default:"5096"`
	Sessions    int `hcl:"Sessions,optional" default:"5096"`
	LinkEmbeds  int `hcl:"LinkEmbeds,optional" default:"65536"`
	InstantView int `hcl:"InstantView,optional" default:"65536"`
}

type MessagesPolicy struct {
	MaximumGetAmount int `hcl:"MaximumGetAmount,optional" default:"50"`
}

type APIPolicy struct {
	Messages MessagesPolicy `hcl:"Messages,block"`
}

type FederationPolicy struct {
	NonceLength                       int `hcl:"NonceLength,optional" default:"32"`
	GuildLeaveNotificationQueueLength int `hcl:"GuildLeaveNotificationQueueLength,optional" default:"64"`
}

type ServerPolicies struct {
	EnablePasswordResetForm bool             `hcl:"EnablePasswordResetForm,optional" default:"false"`
	Avatar                  AvatarPolicy     `hcl:"Avatar,block"`
	Username                UsernamePolicy   `hcl:"Username,block"`
	Password                PasswordPolicy   `hcl:"Password,block"`
	Attachments             AttachmentPolicy `hcl:"Attachments,block"`
	Debug                   DebugPolicy      `hcl:"Debug,block"`
	Sessions                SessionPolicy    `hcl:"Sessions,block"`
	MaximumCacheSizes       CachePolicy      `hcl:"Caches,block"`
	APIs                    APIPolicy        `hcl:"APIs,block"`
	Federation              FederationPolicy `hcl:"Federation,block"`
}

type ServerConf struct {
	Host           string         `hcl:"Host,optional" default:"0.0.0.0"`
	Port           int            `hcl:"Port,optional" default:"2289"`
	PrivateKeyPath string         `hcl:"PrivateKeyPath,optional" default:"harmony-key.pem"`
	PublicKeyPath  string         `hcl:"PublicKeyPath,optional" default:"harmony-key.pub"`
	StorageBackend string         `hcl:"StorageBackend,optional" default:"PureFlatfile"`
	SnowflakeStart int64          `hcl:"SnowflakeStart,optional" default:"0"`
	UseCORS        bool           `hcl:"UseCORS,optional" default:"true"`
	UseTLS         bool           `hcl:"UseTLS,optional" default:"false"`
	TLSCert        string         `hcl:"TLSCert,optional"`
	TLSKey         string         `hcl:"TLSKey,optional"`
	Policies       ServerPolicies `hcl:"Policies,block"`
}

type DBConf struct {
	Host     string `hcl:"Host,optional" default:"127.0.0.1"`
	Username string `hcl:"Username"`
	Password string `hcl:"Password"`
	Port     int    `hcl:"Port,optional" default:"5432"`
	SSL      bool   `hcl:"SSL,optional" default:"false"`
	Name     string `hcl:"Name,optional" default:"harmony"`
	Backend  string `hcl:"Backend,optional" default:"sqlite"`
	Filename string `hcl:"Filename,optional" default:"data.db"`
}

type FlatfileConf struct {
	MediaPath string `hcl:"MediaPath,optional" default:"flatfile"`
}

type SentryConf struct {
	DSN               string `hcl:"DSN,optional"`
	AttachStacktraces bool   `hcl:"AttachStacktraces,optional" default:"true"`
	Enabled           bool   `hcl:"Enabled,optional" default:"false"`
}

type Config struct {
	Server   ServerConf   `hcl:"Server,block"`
	Database DBConf       `hcl:"Database,block"`
	Flatfile FlatfileConf `hcl:"Flatfile,block"`
	Sentry   SentryConf   `hcl:"Sentry,block"`
}

// Load reads a config file (JSON format)
func Load() (*Config, error) {
	var config Config
	defaults.MustSet(&config)

	if _, err := os.Stat("config.hcl"); os.IsNotExist(err) {
		file := hclwrite.NewFile()
		gohcl.EncodeIntoBody(&config, file.Body())

		err = ioutil.WriteFile("config.hcl", file.Bytes(), 0755)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("A default configuration has been written to 'config.hcl'. Edit it as appropriate and then restart Legato.")
		os.Exit(0)
	}

	if err := tracerr.Wrap(hclsimple.DecodeFile("config.hcl", nil, &config)); err != nil {
		return nil, err
	}

	return &config, nil
}

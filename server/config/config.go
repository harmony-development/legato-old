package config

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/creasty/defaults"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/ztrue/tracerr"
)

type Config struct {
	Server struct {
		Host           string `hcl:"Host,optional" default:"0.0.0.0"`
		Port           int    `hcl:"Port,optional" default:"2289"`
		PrivateKeyPath string `hcl:"PrivateKeyPath,optional" default:"harmony-key.pem"`
		PublicKeyPath  string `hcl:"PublicKeyPath,optional" default:"harmony-key.pub"`
		StorageBackend string `hcl:"StorageBackend,optional" default:"PureFlatfile"`
		SnowflakeStart int64  `hcl:"SnowflakeStart,optional" default:"0"`
		UseCORS        bool   `hcl:"UseCORS,optional" default:"true"`

		Policies struct {
			Avatar struct {
				Width   int  `hcl:"Width,optional" default:"256"`
				Height  int  `hcl:"Height,optional" default:"256"`
				Quality int  `hcl:"Quality,optional" default:"50"`
				Crop    bool `hcl:"Crop,optional" default:"true"`
			} `hcl:"Avatar,block"`
			Username struct {
				MinLength int `hcl:"MinLength,optional" default:"2"`
				MaxLength int `hcl:"MaxLength,optional" default:"32"`
			} `hcl:"Username,block"`
			Password struct {
				MinLength  int `hcl:"MinLength,optional" default:"5"`
				MaxLength  int `hcl:"MaxLength,optional" default:"256"`
				MinLower   int `hcl:"MinLower,optional" default:"1"`
				MinUpper   int `hcl:"MinUpper,optional" default:"1"`
				MinNumbers int `hcl:"MinNumbers,optional" default:"1"`
				MinSymbols int `hcl:"MinSymbols,optional" default:"0"`
			} `hcl:"Password,block"`
			Attachments struct {
				MaximumCount int `hcl:"MaximumCount,optional" default:"10"`
			} `hcl:"Attachments,block"`
			Debug struct {
				LogErrors                  bool `hcl:"LogErrors,optional" default:"true"`
				LogRequests                bool `hcl:"LogRequests,optional" default:"true"`
				RespondWithErrors          bool `hcl:"RespondWithErrors,optional" default:"false"`
				ResponseErrorsIncludeTrace bool `hcl:"ResponseErrorsIncludeTrace" default:"true"`
				VerboseStreamHandling      bool `hcl:"VerboseStreamHandling" default:"false"`
			} `hcl:"Debug,block"`
			Sessions struct {
				Duration time.Duration `hcl:"Duration,optional" default:"172800000000000"`
			} `hcl:"Sessions,block"`
			MaximumCacheSizes struct {
				Owner       int `hcl:"Owner,optional" default:"5096"`
				Sessions    int `hcl:"Sessions,optional" default:"5096"`
				LinkEmbeds  int `hcl:"LinkEmbeds,optional" default:"65536"`
				InstantView int `hcl:"InstantView,optional" default:"65536"`
			} `hcl:"Caches,block"`
			APIs struct {
				Messages struct {
					MaximumGetAmount int `hcl:"MaximumGetAmount,optional" default:"50"`
				} `hcl:"Messages,block"`
			} `hcl:"APIs,block"`
			Federation struct {
				NonceLength                       int `hcl:"NonceLength,optional" default:"32"`
				GuildLeaveNotificationQueueLength int `hcl:"GuildLeaveNotificationQueueLength,optional" default:"64"`
			} `hcl:"Federation,block"`
		} `hcl:"Policies,block"`
	} `hcl:"Server,block"`
	Database struct {
		Host     string `hcl:"Host,optional" default:"127.0.0.1"`
		Username string `hcl:"Username"`
		Password string `hcl:"Password"`
		Port     int    `hcl:"Port,optional" default:"5432"`
		SSL      bool   `hcl:"SSL,optional" default:"false"`
		Name     string `hcl:"Name,optional" default:"harmony"`
	} `hcl:"Database,block"`
	Flatfile struct {
		MediaPath string `hcl:"MediaPath,optional" default:"flatfile"`
	} `hcl:"Flatfile,block"`
	Sentry struct {
		DSN               string `hcl:"DSN,optional"`
		AttachStacktraces bool   `hcl:"AttachStacktraces,optional" default:"true"`
		Enabled           bool   `hcl:"Enabled,optional" default:"false"`
	} `hcl:"Sentry,block"`
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

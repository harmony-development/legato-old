package config_test

import (
	"strings"
	"testing"

	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/test"
	"github.com/stretchr/testify/assert"
)

func TestConfigLoad(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	cfg := `
server:
  port: 1000
  cassandra:
    numretries: 5`

	conf, err := config.LoadConfig(strings.NewReader(cfg))

	a.Nil(err)
	a.Equal(1000, conf.Server.Port)
	a.Equal(5, conf.Server.Cassandra.NumRetries)
}

func TestConfigDefaults(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	cfg := ``

	conf, err := config.LoadConfig(strings.NewReader(cfg))

	a.Nil(err)
	a.Equal(2289, conf.Server.Port)
}

func TestInvalidConfig(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	_, err := config.LoadConfig(strings.NewReader(`this isn't yaml`))

	a.NotNil(err)
}

// This test kind of feels pointless tbh.
func TestInvalidReader(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	_, err := config.LoadConfig(test.ErrReader(0))

	a.EqualError(err, "test error")
}

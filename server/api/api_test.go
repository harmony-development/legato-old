package api_test

import (
	"testing"

	"github.com/harmony-development/legato/server/api"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	i := api.New()
	assert.NotNil(t, i.Echo)
}

func TestStartServes(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	i := api.New()
	serviceRunning := make(chan interface{})
	serviceDone := make(chan interface{})

	var err error

	go func() {
		close(serviceRunning)

		err = i.Start(":0")

		defer close(serviceDone)
	}()

	<-serviceRunning
	a.NoError(i.Shutdown())
	a.NoError(err)
}

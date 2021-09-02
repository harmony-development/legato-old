package server

import (
	"strconv"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/build"
	"github.com/harmony-development/legato/config"
)

type Server struct {
	*fiber.App
	l log.Interface
	c *config.Config
}

var startupMessage = `	
   __                  __      
  / /___  ___ _ ___ _ / /_ ___ 
 / // -_)/ _ ` + "`" + `// _ ` + "`" + `// __// _ \
/_/ \__/ \_, / \_,_/ \__/ \___/
        /___/ Version %s
`

func New(l log.Interface, cfg *config.Config) *Server {
	s := &Server{
		App: fiber.New(fiber.Config{
			AppName:               "legato",
			DisableStartupMessage: true,
		}),
		c: cfg,
		l: l,
	}
	api.Setup(l, s.App)
	return s
}

func (s *Server) Start() {
	s.l.Infof(startupMessage, color.GreenString(build.GitCommit))
	s.l.Infof("Listening on %s:%d", s.c.Address, s.c.Port)
	s.l.WithError(
		s.Listen(s.c.Address + ":" + strconv.Itoa(s.c.Port)),
	).Fatal("Fatal error occured")
}

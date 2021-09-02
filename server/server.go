package server

import (
	"fmt"
	"strconv"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/build"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/logger"
)

type Server struct {
	*fiber.App
	l log.Interface
	c *config.Config
}

var startupMessage = `Version %s
   __                  __      
  / /___  ___ _ ___ _ / /_ ___ 
 / // -_)/ _ ` + "`" + `// _ ` + "`" + `// __// _ \
/_/ \__/ \_, / \_,_/ \__/ \___/
        /___/ Commit %s
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

func (s *Server) printStartup() {
	display := logger.Indent(
		log.InfoLevel,
		startupMessage,
		color.HiMagentaString(fmt.Sprintf("Listening on %s:%d", s.c.Address, s.c.Port)),
	)
	versionString := color.GreenString(build.Version)
	gitString := color.GreenString(
		fmt.Sprintf("%.7s", build.GitCommit),
	)
	s.l.Infof(display, versionString, gitString)
}

func (s *Server) Start() {
	s.printStartup()
	fmt.Print("   >  ")
	s.l.WithError(
		s.Listen(s.c.Address + ":" + strconv.Itoa(s.c.Port)),
	).Fatal("Fatal error occured")
}

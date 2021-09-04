package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/harmony-development/legato/build"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/logger"
	"github.com/harmony-development/legato/server/api"
	authv1impl "github.com/harmony-development/legato/server/api/authv1"
	chatv1impl "github.com/harmony-development/legato/server/api/chatv1"
	"github.com/harmony-development/legato/server/key"
)

type Server struct {
	*fiber.App
	l          log.Interface
	c          *config.Config
	keyManager key.KeyManager
}

var startupMessage = `Version %s
   __                  __      
  / /___  ___ _ ___ _ / /_ ___ 
 / // -_)/ _ ` + "`" + `// _ ` + "`" + `// __// _ \
/_/ \__/ \_, / \_,_/ \__/ \___/
        /___/ Commit %s
`

func New(l *logger.Logger, cfg *config.Config, db db.DB) (*Server, error) {
	s := &Server{
		App: fiber.New(fiber.Config{
			AppName:               "legato",
			DisableStartupMessage: true,
		}),
		c: cfg,
		l: l,
	}
	s.setupMiddleware()

	keyManager, err := key.NewEd25519KeyManagerFromFile(s.c.PrivateKeyPath, s.c.PublicKeyPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := key.WriteEd25519KeysToFile(s.c.PrivateKeyPath, s.c.PublicKeyPath); err != nil {
				return nil, err
			}
			keyManager, err = key.NewEd25519KeyManagerFromFile(s.c.PrivateKeyPath, s.c.PublicKeyPath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	register := api.Setup(l, s.App)

	register(authv1.NewAuthServiceHandler(
		authv1impl.New(
			keyManager,
			db,
		),
	))
	register(chatv1.NewChatServiceHandler(chatv1impl.ChatV1{}))

	return s, nil
}

func (s *Server) setupMiddleware() {
	s.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} |${green}${method}${white}|  ${path}  ↑${bytesSent} bytes  ↓${bytesReceived} bytes ${reqHeader:Authorization}\n",
	}))
}

func (s *Server) printStartup() {
	listenText := color.HiMagentaString(fmt.Sprintf("Listening on %s:%d", s.c.Address, s.c.Port))
	display := logger.Indent(
		log.InfoLevel,
		startupMessage,
		&listenText,
	)
	versionString := color.GreenString(build.Version)
	gitString := color.GreenString(
		fmt.Sprintf("%.7s", build.GitCommit),
	)
	s.l.Infof(display, versionString, gitString)
	fmt.Print("   >  ")
}

func (s *Server) Start() {
	s.printStartup()

	s.l.WithError(
		s.Listen(s.c.Address + ":" + strconv.Itoa(s.c.Port)),
	).Fatal("Fatal error occured")
}

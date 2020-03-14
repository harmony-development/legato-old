package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"harmony-auth-server/rest/hm"
	v1 "harmony-auth-server/rest/v1"
	"time"
)

func Setup(g echo.Group) {
	go hm.CleanupRoutine()
	apiV1(g)
}

func apiV1(g echo.Group) {
	r := g.Group("/v1")
	r.Use(hm.WithHarmony)
	r.Use(middleware.CORS())
	r.POST("/login*", hm.WithRateLimit(v1.Login, 5 * time.Second, 5))
	r.POST("/register", hm.WithRateLimit(v1.Register, 15 * time.Second, 8))
	r.POST("/getuser*", hm.WithRateLimit(v1.GetUser, 5 * time.Second, 5))
	r.POST("/usernameupdate**", hm.WithRateLimit(v1.UsernameUpdate, 2 * time.Second, 5))
	r.POST("/auth*", hm.WithRateLimit(v1.Authenticate, 2 * time.Second, 5))
	r.POST("/listservers*", hm.WithRateLimit(v1.ListServers, 5 * time.Second, 5))
}
package rest

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/rest/hm"
	v1 "harmony-auth-server/rest/v1"
	"time"
)

func Setup(g *echo.Group) {
	go hm.CleanupRoutine()
	apiV1(g)
}

func apiV1(g *echo.Group) {
	g.Use(hm.WithHarmony)
	g.POST("/login*", hm.WithRateLimit(v1.Login, 5 * time.Second, 5))
	g.POST("/getuser*", hm.WithRateLimit(v1.GetUser, 5 * time.Second, 5))
	g.POST("/usernameupdate**", hm.WithRateLimit(v1.UsernameUpdate, 2 * time.Second, 5))
	g.POST("/auth*", hm.WithRateLimit(v1.Authenticate, 2 * time.Second, 5))
	g.POST("/listservers*", hm.WithRateLimit(v1.ListServers, 5 * time.Second, 5))
}
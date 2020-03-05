package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "harmony-server/rest/v1"
	"time"
)

func SetupREST(g echo.Group) {
	apiV1(g)
}

func apiV1(g echo.Group) {
	r := g.Group("/v1")
	r.Use(middleware.CORS())
	r.POST("/login*", v1.WithRateLimit(v1.Login, 5 * time.Second, 1))
	r.POST("/register*", v1.WithRateLimit(v1.Register, 10 * time.Minute, 10))
	r.POST("/getguilds*", v1.WithRateLimit(v1.GetGuilds, 5 * time.Second, 3))
	r.POST("/getchannels*", v1.WithRateLimit(v1.GetChannels, 500 * time.Millisecond, 5))
	r.POST("/avatarupdate*", v1.WithRateLimit(v1.AvatarUpdate, 3 * time.Second, 1))
	r.POST("/updateguildpicture*", v1.WithRateLimit(v1.UpdateGuildPicture, 3 * time.Second, 1))
	r.POST("/message*", v1.WithRateLimit(v1.Message, 500 * time.Millisecond, 20))
	r.POST("/createguild*", v1.WithRateLimit(v1.CreateGuild, 20 * time.Second, 3))
	r.POST("/addchannel*", v1.WithRateLimit(v1.AddChannel, 1 * time.Second, 3))
	r.POST("/deletechannel*", v1.WithRateLimit(v1.DeleteChannel, 1 * time.Second, 3))
	r.POST("/deleteguild*", v1.WithRateLimit(v1.DeleteGuild, 5 * time.Second, 5))
	r.POST("/deletemessage*", v1.WithRateLimit(v1.DeleteGuild, 1 * time.Second, 10))
	r.POST("/deleteinvite*", v1.WithRateLimit(v1.DeleteInvite, 1 * time.Second, 5))
	r.POST("/createinvite*", v1.WithRateLimit(v1.CreateInvite, 1 * time.Second, 5))
	r.POST("/getinvites*", v1.WithRateLimit(v1.GetInvites, 2 * time.Second, 4))
}
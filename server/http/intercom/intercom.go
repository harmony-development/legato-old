package intercom

import (
	"harmony-server/server/config"
	"harmony-server/server/logger"
)

type IntercomManager struct {
	Logger               logger.ILogger
	GuildLeaveNotifQueue chan string
}

type Dependencies struct {
	Logger logger.ILogger
	Config config.Config
}

func New(deps Dependencies) IntercomManager {
	m := IntercomManager{}
	m.Logger = deps.Logger
	m.GuildLeaveNotifQueue = make(chan string, deps.Config.Server.GuildLeaveNotificationQueueLength)
	for i := 1; i < deps.Config.Server.GuildLeaveNotificationQueueLength; i++ {
		go m.GuildLeaveNotifyRoutine()
	}
	return m
}

func (m IntercomManager) GuildLeaveNotifyRoutine() {
	for {
		r := <-m.GuildLeaveNotifQueue
		m.NotifyLeaveGuild(r)
	}
}

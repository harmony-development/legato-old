package v1

func (s *GuildState) SubAdd(guild, user uint64) {
	g := s.subs
	if _, ok := g[GuildID(guild)]; !ok {
		g[GuildID(guild)] = map[UserID]struct{}{}
	}
	g[GuildID(guild)][UserID(user)] = struct{}{}
}

func (s *GuildState) SubRemove(guild, user uint64) {
	g := s.subs
	if val, ok := g[GuildID(guild)]; ok {
		delete(val, UserID(user))
	}
}

func (s *GuildState) SubRemoveUser(user uint64) {
	g := s.subs
	if val, ok := s.guildEvents[UserID(user)]; ok {
		for key := range val {
			if val, ok := g[key]; ok {
				delete(val, UserID(user))
			}
		}
	}
}

func (s *GuildState) SubRemoveUserFromGuild(user, guild uint64) {
	g := s.subs
	if _, ok := g[GuildID(guild)]; ok {
		delete(g[GuildID(guild)], UserID(user))
	}
}

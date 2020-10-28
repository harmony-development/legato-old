package integrated

func (s *GuildState) subAdd(guild, user uint64) {
	g := s.subs
	if _, ok := g[_guildID(guild)]; !ok {
		g[_guildID(guild)] = map[_userID]struct{}{}
	}
	g[_guildID(guild)][_userID(user)] = struct{}{}
}

func (s *GuildState) subRemove(guild, user uint64) {
	g := s.subs
	if val, ok := g[_guildID(guild)]; ok {
		delete(val, _userID(user))
	}
}

func (s *GuildState) subRemoveUser(user uint64) {
	g := s.subs
	if val, ok := s.guildEvents[_userID(user)]; ok {
		for key := range val {
			if val, ok := g[key]; ok {
				delete(val, _userID(user))
			}
		}
	}
}

func (s *GuildState) subRemoveUserFromGuild(user, guild uint64) {
	g := s.subs
	if _, ok := g[_guildID(guild)]; ok {
		delete(g[_guildID(guild)], _userID(user))
	}
}

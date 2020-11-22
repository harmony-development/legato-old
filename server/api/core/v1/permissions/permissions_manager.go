package permissions

import (
	"encoding/json"

	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/db"
	lru "github.com/hashicorp/golang-lru"
)

// Manager manages permissions
type Manager struct {
	states *lru.Cache
	db     db.IHarmonyDB
}

// NewManager creates a new permissions manager
func NewManager(db db.IHarmonyDB) Manager {
	man := Manager{
		db: db,
	}
	cache, err := lru.NewWithEvict(50_000, func(key, value interface{}) {
		man.saveGuild(key.(uint64), value.(*GuildState))
	})
	if err != nil {
		panic(err)
	}
	man.states = cache
	return man
}

func (p *Manager) saveGuild(guild uint64, data *GuildState) {
	item, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = p.db.SetGuildPermissions(guild, item)
	if err != nil {
		panic(err)
	}
}

func (p *Manager) obtainGuild(guild uint64) *GuildState {
	data, err := p.db.GetGuildPermissions(guild)
	if err != nil {
		panic(err)
	}
	gs := new(GuildState)
	gs.Categories = make(map[ChannelID]ChannelID)
	gs.Roles = make(map[RoleID][]PermissionNode)
	gs.Channels = make(map[ChannelID]map[RoleID][]PermissionNode)
	err = json.Unmarshal(data, gs)
	if err != nil {
		panic(err)
	}

	return gs
}

// Check checks whether a user with the given roles has a permission in a given channel
func (p *Manager) Check(permission string, userRoles []uint64, inGuild uint64, inChannel uint64) bool {
	if !p.states.Contains(inGuild) {
		p.states.Add(inGuild, p.obtainGuild(inGuild))
	}
	data, _ := p.states.Get(inGuild)
	state := data.(GuildState)
	return state.Check(permission, userRoles, ChannelID(inChannel))
}

func (p *Manager) SetPermissions(permissions []*corev1.Permission, forGuild, forChannel, forRole uint64) error {
	var guild *GuildState

	if !p.states.Contains(forGuild) {
		guild = p.obtainGuild(forGuild)
	} else {
		intf, _ := p.states.Get(forGuild)
		guild = intf.(*GuildState)
	}

	var nodes []PermissionNode
	for _, perm := range permissions {
		node := PermissionNode{}

		var err error
		node.Glob, err = TryGlob(perm.Matches)

		if err != nil {
			return err
		}

		if perm.Mode == corev1.Permission_Allow {
			node.Mode = Allow
		} else {
			node.Mode = Deny
		}

		nodes = append(nodes, node)
	}

	if forChannel == 0 {
		guild.Roles[RoleID(forRole)] = nodes
	} else {
		if _, ok := guild.Channels[ChannelID(forChannel)]; !ok {
			guild.Channels[ChannelID(forChannel)] = make(map[RoleID][]PermissionNode)
		}
		guild.Channels[ChannelID(forChannel)][RoleID(forRole)] = nodes
	}

	if !p.states.Contains(forGuild) {
		p.saveGuild(forGuild, guild)
	}

	return nil
}

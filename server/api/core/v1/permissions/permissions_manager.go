package permissions

import (
	"github.com/alecthomas/repr"
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
func NewManager(db db.IHarmonyDB) *Manager {
	man := &Manager{
		db: db,
	}
	cache, err := lru.New(50_000)
	if err != nil {
		panic(err)
	}
	man.states = cache
	return man
}

var zeroVal = PermissionNode{}

func (p *Manager) saveGuild(guild uint64, data *GuildState) {
	for channel, cdata := range data.Channels {
		for role, rdata := range cdata {
			if err := p.db.SetPermissions(guild, uint64(channel), uint64(role), func() (ret []db.PermissionsNode) {
				for _, perm := range rdata {
					if perm == zeroVal {
						continue
					}
					ret = append(ret, db.PermissionsNode{
						Node:  perm.Glob.s,
						Allow: perm.Mode == Allow,
					})
				}
				return
			}()); err != nil {
				panic(err)
			}
		}
	}

	for role, rdata := range data.Roles {
		if err := p.db.SetPermissions(guild, 0, uint64(role), func() (ret []db.PermissionsNode) {
			for _, perm := range rdata {
				if perm == zeroVal {
					continue
				}
				ret = append(ret, db.PermissionsNode{
					Node:  perm.Glob.s,
					Allow: perm.Mode == Allow,
				})
			}
			return
		}()); err != nil {
			panic(err)
		}
	}
}

func (p *Manager) obtainGuild(guild uint64) *GuildState {
	data, err := p.db.GetPermissionsData(guild)

	repr.Println(data)

	if err != nil {
		panic(err)
	}

	gs := new(GuildState)
	gs.Categories = make(map[ChannelID]ChannelID)
	gs.Roles = make(map[RoleID][]PermissionNode)
	gs.Channels = make(map[ChannelID]map[RoleID][]PermissionNode)

	dbToManager := func(nodes []db.PermissionsNode) (ret []PermissionNode) {
		for _, node := range nodes {
			ret = append(ret, PermissionNode{
				Glob: MustGlob(node.Node),
				Mode: func() Mode {
					if node.Allow {
						return Allow
					}
					return Deny
				}(),
			})
		}
		return
	}

	for id, category := range data.Categories {
		for _, channel := range category {
			gs.Categories[ChannelID(channel)] = ChannelID(id)
		}
	}
	for channelID, channel := range data.Channels {
		gs.Channels[ChannelID(channelID)] = make(map[RoleID][]PermissionNode)
		for roleID, role := range channel {
			gs.Channels[ChannelID(channelID)][RoleID(roleID)] = dbToManager(role)
		}
	}
	for roleID, role := range data.Roles {
		gs.Roles[RoleID(roleID)] = dbToManager(role)
	}

	return gs
}

// Check checks whether a user with the given roles has a permission in a given channel
func (p *Manager) Check(permission string, userRoles []uint64, inGuild uint64, inChannel uint64) bool {
	if !p.states.Contains(inGuild) {
		p.states.Add(inGuild, p.obtainGuild(inGuild))
	}
	data, _ := p.states.Get(inGuild)
	state := data.(*GuildState)
	return state.Check(permission, userRoles, ChannelID(inChannel))
}

func (p *Manager) ensureGuild(guildID uint64) *GuildState {
	if !p.states.Contains(guildID) {
		p.states.Add(guildID, p.obtainGuild(guildID))
	}

	val, _ := p.states.Get(guildID)

	return val.(*GuildState)
}

func (p *Manager) GetPermissions(forGuild, forChannel, forRole uint64) (ret []*corev1.Permission) {
	guild := p.ensureGuild(forGuild)

	if forChannel == 0 {
		data := guild.Roles[RoleID(forRole)]
		for _, node := range data {
			ret = append(ret, &corev1.Permission{
				Matches: node.Glob.s,
				Mode: func() corev1.Permission_Mode {
					if node.Mode == Allow {
						return corev1.Permission_Allow
					}
					return corev1.Permission_Deny
				}(),
			})
		}
	} else {
		if _, ok := guild.Channels[ChannelID(forChannel)]; !ok {
			return
		}
		data := guild.Channels[ChannelID(forChannel)][RoleID(forRole)]
		for _, node := range data {
			ret = append(ret, &corev1.Permission{
				Matches: node.Glob.s,
				Mode: func() corev1.Permission_Mode {
					if node.Mode == Allow {
						return corev1.Permission_Allow
					}
					return corev1.Permission_Deny
				}(),
			})
		}
	}

	repr.Println(ret)

	return
}

func (p *Manager) SetPermissions(permissions []*corev1.Permission, forGuild, forChannel, forRole uint64) error {
	guild := p.ensureGuild(forGuild)

	repr.Println(permissions)

	if !p.states.Contains(forGuild) {
		guild = p.obtainGuild(forGuild)
	} else {
		intf, _ := p.states.Get(forGuild)
		guild = intf.(*GuildState)
	}

	var nodes []PermissionNode
	for _, perm := range permissions {
		node := PermissionNode{}

		if perm.Matches == "" {
			continue
		}

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

	go p.saveGuild(forGuild, guild)

	return nil
}

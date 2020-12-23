package permissions

import (
	"encoding/json"

	"github.com/gobwas/glob"
)

// Mode determines whether a permission will glob or not
type Mode int

// RoleID is the ID of a role
type RoleID uint64

// ChannelID is the ID of a channel
type ChannelID uint64

const (
	// Allow permission
	Allow Mode = iota
	// Deny permission
	Deny
)

// Everyone has this role. You don't need to explicitly specify it in GuildState.Check.
const Everyone RoleID = 0

// PermissionGlob is a glob type
type PermissionGlob struct {
	g glob.Glob
	s string
}

// MustGlob will panic if the glob isn't valid
func MustGlob(s string) (p PermissionGlob) {
	p.g = glob.MustCompile(s)
	p.s = s
	return
}

// TryGlob tries to create a PermissionGlob
func TryGlob(s string) (p PermissionGlob, err error) {
	p.g, err = glob.Compile(s)
	p.s = s
	return
}

// Match returns whether or not the given string matches
func (p *PermissionGlob) Match(str string) bool {
	return p.g.Match(str)
}

// UnmarshalJSON impl
func (p *PermissionGlob) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	*p, err = TryGlob(str)
	return err
}

// MarshalJSON impl
func (p *PermissionGlob) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.s)
}

// PermissionNode holds a glob and what it says about permissions
type PermissionNode struct {
	Glob PermissionGlob
	Mode
}

// GuildState represents the permissions used by a guild
type GuildState struct {
	Roles      map[RoleID][]PermissionNode
	Categories map[ChannelID]ChannelID
	Channels   map[ChannelID]map[RoleID][]PermissionNode
}

// Check whether the user with the given roles has the permission to do something
// in the given channel ID. userRoles should have the most important roles first
// and the least important ones last.
func (g GuildState) Check(permission string, userRoles []uint64, in ChannelID) bool {
	userRoles = append(userRoles, uint64(Everyone))

	if in != 0 {
		if channelData, ok := g.Channels[in]; ok {
			for _, role := range userRoles {
				nodes, ok := channelData[RoleID(role)]
				_ = ok
				for _, node := range nodes {
					if node.Glob.Match(permission) {
						return node.Mode == Allow
					}
				}
			}
		}

		if category, ok := g.Categories[in]; ok {
			if channelData, ok := g.Channels[category]; ok {
				for _, role := range userRoles {
					nodes, ok := channelData[RoleID(role)]
					_ = ok
					for _, node := range nodes {
						if node.Glob.Match(permission) {
							return node.Mode == Allow
						}
					}
				}
			}
		}
	}

	for _, role := range userRoles {
		nodes, ok := g.Roles[RoleID(role)]
		_ = ok
		for _, node := range nodes {
			if node.Glob.Match(permission) {
				return node.Mode == Allow
			}
		}
	}

	return false
}

package types

import (
	"fmt"
	"strings"
)

type PermissionsNode struct {
	Node  string
	Allow bool
}

func (p *PermissionsNode) Deserialize(s string) (ok bool) {
	trimmed := strings.Split(strings.TrimSuffix(strings.TrimPrefix(s, "("), ")"), ",")

	if len(trimmed) != 3 {
		return false
	}

	if trimmed[2] == "t" {
		p.Allow = true
	} else {
		p.Allow = false
	}

	p.Node = trimmed[1]

	return true
}

func (p PermissionsNode) Serialize() string {
	node := "f"
	if p.Allow {
		node = "t"
	}
	return fmt.Sprintf("(%s,%s)", p.Node, node)
}

type PermissionsData struct {
	Roles      map[uint64][]PermissionsNode
	Categories map[uint64][]uint64
	Channels   map[uint64]map[uint64][]PermissionsNode
}

type UserData struct {
	UserID   uint64
	Email    string
	Username string
	Avatar   string
	Status   int16
	IsBot    bool
	Password []byte
}

type GuildData struct {
	ID      uint64
	Owner   uint64
	Name    string
	Picture string
}

type RoleData struct{}

type GuildListEntryData struct {
	ID   uint64
	Host string
}

type MessageData struct {
	GuildID   uint64
	ChannelID uint64
	Actions   []byte
}

type EmotePackData struct {
}

type EmoteData struct {
}

type FileData struct {
}

type InviteData struct {
	ID           string
	PossibleUses int32
	Uses         int32
}

type ChannelData struct {
	ID       uint64
	Name     string
	Metadata []byte
}

type ChannelKind uint64

const (
	ChannelKindText ChannelKind = iota
	ChannelKindCategory
	ChannelKindVoice
)

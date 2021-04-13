package types

import (
	"fmt"
	"strings"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"google.golang.org/protobuf/types/known/timestamppb"
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

type RoleData struct {
	ID       uint64
	Name     string
	Position string
	Color    int
	Hoist    bool
	Pingable bool
}

type GuildListEntryData struct {
	ID       uint64
	Host     string
	Position string
}

type EmotePackData struct {
	OwnerID uint64
	PackID  uint64
	Name    string
}

type EmoteData struct {
	Name    string
	ImageID string
}

type FileData struct {
	FileID      string
	ContentType string
	Name        string
	Size        int
}

type InviteData struct {
	ID           string
	PossibleUses int32
	Uses         int32
}

type ChannelData struct {
	ID       uint64
	Name     string
	Position string
	Kind     uint64
	Metadata *harmonytypesv1.Metadata
}

type MessageOverride struct {
	Username string
	Avatar   string
	Reason   string
}

type MessageData struct {
	Metadata  *harmonytypesv1.Metadata
	Overrides *harmonytypesv1.Override
	GuildId   uint64
	ChannelId uint64
	MessageId uint64
	AuthorId  uint64
	CreatedAt *timestamppb.Timestamp
	EditedAt  *timestamppb.Timestamp
	InReplyTo uint64
	Content   *harmonytypesv1.Content
}

func (v *MessageData) ToV1() *harmonytypesv1.Message {
	return &harmonytypesv1.Message{
		Metadata:  v.Metadata,
		Overrides: v.Overrides,
		GuildId:   v.GuildId,
		MessageId: v.MessageId,
		AuthorId:  v.AuthorId,
		CreatedAt: v.CreatedAt,
		EditedAt:  v.EditedAt,
		InReplyTo: v.InReplyTo,
		Content:   v.Content,
	}
}

func ManyMessageDataInto(msgs []*MessageData) (r []*harmonytypesv1.Message) {
	for _, it := range msgs {
		r = append(r, it.ToV1())
	}
	return
}

func Into(msg *entgen.Message) *MessageData {
	return &MessageData{
		GuildId:   msg.Edges.Channel.Edges.Guild.ID,
		ChannelId: msg.Edges.Channel.ID,
		MessageId: msg.ID,
		AuthorId:  msg.Edges.User.ID,
		CreatedAt: timestamppb.New(msg.Createdat),
		EditedAt:  timestamppb.New(msg.Editedat),
		InReplyTo: func() uint64 {
			if msg.Edges.Parent != nil {
				return msg.Edges.Parent.ID
			}
			return 0
		}(),
		Content: msg.Content,
	}
}

func IntoMany(msgs []*entgen.Message) (out []*MessageData) {
	for _, it := range msgs {
		out = append(out, Into(it))
	}
	return
}

type ChannelKind uint64

const (
	ChannelKindText ChannelKind = iota
	ChannelKindCategory
	ChannelKindVoice
)

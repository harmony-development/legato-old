package permissions

import (
	"encoding/json"
	"testing"
)

const (
	BonkRole             RoleID = 1
	MoyaiRole            RoleID = 2
	EpiclyMutedGamerRole RoleID = 3
)

const (
	GeneralChannel              ChannelID = 1
	ModerationCategory          ChannelID = 2
	ModerationDiscussionChannel ChannelID = 3
	ModerationBotChannel        ChannelID = 4
	SecretChannel               ChannelID = 5
)

var PermissionData = GuildState{
	Roles: map[RoleID][]PermissionNode{
		BonkRole: {
			{
				Glob: MustGlob("*"),
				Mode: Allow,
			},
		},
		MoyaiRole: {
			{
				Glob: MustGlob("moderation.*"),
				Mode: Allow,
			},
		},
		EpiclyMutedGamerRole: {
			{
				Glob: MustGlob("chat.sendMessages"),
				Mode: Deny,
			},
		},
		Everyone: {
			{
				Glob: MustGlob("chat.*"),
				Mode: Allow,
			},
		},
	},
	Categories: map[ChannelID]ChannelID{
		ModerationBotChannel:        ModerationCategory,
		ModerationDiscussionChannel: ModerationCategory,
		SecretChannel:               ModerationCategory,
	},
	Channels: map[ChannelID]map[RoleID][]PermissionNode{
		ModerationCategory: {
			MoyaiRole: {
				{
					Glob: MustGlob("chat.*"),
					Mode: Allow,
				},
			},
			Everyone: {
				{
					Glob: MustGlob("chat.*"),
					Mode: Deny,
				},
			},
		},
		SecretChannel: {
			BonkRole: {
				{
					Glob: MustGlob("*"),
					Mode: Allow,
				},
			},
			Everyone: {
				{
					Glob: MustGlob("*"),
					Mode: Deny,
				},
			},
		},
	},
}

type Check struct {
	Node     string
	Expected bool
	In       ChannelID
}

type TestingData struct {
	Roles []uint64
	Check []Check
}

var Yeet = map[string]TestingData{
	"jan Pontajosi": {
		Roles: []uint64{uint64(BonkRole), uint64(MoyaiRole)},
		Check: []Check{
			{
				Node:     "chat.chat",
				Expected: true,
				In:       SecretChannel,
			},
		},
	},
	"Blusky": {
		Roles: []uint64{uint64(MoyaiRole)},
		Check: []Check{
			{
				Node:     "chat.chat",
				Expected: false,
				In:       SecretChannel,
			},
			{
				Node:     "chat.chat",
				Expected: true,
				In:       ModerationDiscussionChannel,
			},
			{
				Node:     "moderation.ban",
				Expected: true,
				In:       GeneralChannel,
			},
		},
	},
	"Some Random Joe": {
		Roles: []uint64{},
		Check: []Check{
			{
				Node:     "chat.chat",
				Expected: false,
				In:       SecretChannel,
			},
			{
				Node:     "chat.chat",
				Expected: false,
				In:       ModerationDiscussionChannel,
			},
			{
				Node:     "moderation.ban",
				Expected: false,
				In:       GeneralChannel,
			},
		},
	},
}

func TestPermissions(t *testing.T) {
	for name, data := range Yeet {
		t.Logf("Testing user '%s'...", name)

		for _, item := range data.Check {
			if PermissionData.Check(item.Node, data.Roles, item.In) != item.Expected {
				t.FailNow()
			}
		}
	}
}

func TestSerialize(t *testing.T) {
	data, err := json.Marshal(PermissionData)
	if err != nil {
		t.Fatal(err)
	}

	var g GuildState
	err = json.Unmarshal(data, &g)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkPermissions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, data := range Yeet {
			for _, item := range data.Check {
				PermissionData.Check(item.Node, data.Roles, item.In)
			}
		}
	}
}

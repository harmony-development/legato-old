package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/emotepack"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

func (d *DB) CreateEmotePack(userID, packID uint64, packName string) (err error) {
	defer doRecovery(&err)

	d.EmotePack.
		Create().
		SetID(packID).
		SetUserID(userID).
		SetName(packName).
		SaveX(ctx)

	return
}

func (d *DB) IsPackOwner(userID, packID uint64) (isOwner bool, err error) {
	defer doRecovery(&err)
	isOwner = d.EmotePack.
		Query().
		Where(
			emotepack.ID(userID),
			emotepack.HasUserWith(user.ID(userID)),
		).
		ExistX(ctx)
	return
}

func (d *DB) AddEmoteToPack(packID uint64, imageID string, name string) (err error) {
	defer doRecovery(&err)
	d.EmotePack.UpdateOneID(packID).AddEmote(
		d.Emote.
			Create().
			SetID(imageID).
			SetName(name).
			SaveX(ctx),
	).ExecX(ctx)
	return
}

func (d *DB) DeleteEmoteFromPack(packID uint64, emoteID string) (err error) {
	defer doRecovery(&err)
	d.EmotePack.UpdateOneID(packID).RemoveEmoteIDs(emoteID).ExecX(ctx)
	return
}

func (d *DB) DeleteEmotePack(packID uint64) (err error) {
	defer doRecovery(&err)
	d.EmotePack.DeleteOneID(packID).ExecX(ctx)
	return
}

func (d *DB) GetEmotePacks(userID uint64) (packs []*entgen.EmotePack, err error) {
	defer doRecovery(&err)
	packs = d.User.GetX(ctx, userID).QueryEmotepack().WithOwner().WithEmote().AllX(ctx)
	return
}

func (d *DB) GetEmotePackEmotes(packID uint64) (emotes []*entgen.Emote, err error) {
	defer doRecovery(&err)
	emotes = d.EmotePack.GetX(ctx, packID).QueryEmote().AllX(ctx)
	return
}

func (d *DB) DequipEmotePack(userID, packID uint64) (err error) {
	defer doRecovery(&err)
	d.User.GetX(ctx, userID).Update().RemoveEmotepackIDs(packID)
	return
}

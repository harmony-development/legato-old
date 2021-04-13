package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/emotepack"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/types"
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

func (d *DB) GetEmotePacks(userID uint64) (packs []*types.EmotePackData, err error) {
	defer doRecovery(&err)
	data := d.User.GetX(ctx, userID).QueryEmotepack().WithOwner().WithEmote().AllX(ctx)
	packs = make([]*types.EmotePackData, len(data))
	for i, pack := range data {
		packs[i] = &types.EmotePackData{
			Name:    pack.Name,
			PackID:  pack.ID,
			OwnerID: pack.QueryUser().OnlyIDX(ctx),
		}
	}
	return
}

func (d *DB) GetEmotePackEmotes(packID uint64) (emotes []*types.EmoteData, err error) {
	defer doRecovery(&err)
	data := d.EmotePack.GetX(ctx, packID).QueryEmote().AllX(ctx)
	emotes = make([]*types.EmoteData, len(data))
	for i, pack := range data {
		emotes[i] = &types.EmoteData{
			Name:    pack.Name,
			ImageID: pack.ID,
		}
	}
	return
}

func (d *DB) DequipEmotePack(userID, packID uint64) (err error) {
	defer doRecovery(&err)
	d.User.GetX(ctx, userID).Update().RemoveEmotepackIDs(packID)
	return
}

package ent_shared

import (
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/foreignuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/localuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/profile"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/ent/entgen/usermeta"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *DB) AddForeignUser(host string, userID, localUserID uint64, username, avatar string) (err error) {
	defer doRecovery(&err)
	tx := d.TxX()
	tx.ForeignUser.Create().
		SetForeignid(userID).
		SetHost(host).
		SetUser(
			tx.User.Create().
				SetProfile(
					tx.Profile.Create().
						SetAvatar(avatar).
						SetUsername(username).
						SaveX(ctx),
				).
				SaveX(ctx),
		).
		SaveX(ctx)
	if err := tx.Commit(); err != nil {
		chk(tx.Rollback())
	}
	return
}

func (d *DB) AddLocalUser(userID uint64, email, username string, passwordHash []byte) (err error) {
	defer doRecovery(&err)
	tx := d.TxX()
	tx.LocalUser.
		Create().
		SetEmail(email).
		SetPassword(passwordHash).
		SetUser(
			d.User.
				Create().
				SetID(userID).
				SetProfile(
					d.Profile.
						Create().
						SetUsername(username).
						SaveX(ctx),
				).
				SaveX(ctx),
		).
		SaveX(ctx)
	if err := tx.Commit(); err != nil {
		chk(tx.Rollback())
	}
	return
}

func (d *DB) LocalUserIDToForeignUserID(id uint64) (ret uint64, host string, err error) {
	defer doRecovery(&err)

	data := d.User.GetX(ctx, id).QueryForeignUser().OnlyX(ctx)

	return data.Foreignid, data.Host, nil
}

func (d *DB) EmailExists(email string) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.LocalUser.Query().Where(localuser.Email(email)).ExistX(ctx)
	return
}

func (d *DB) GetAvatar(userID uint64) (avatar *string, err error) {
	defer doRecovery(&err)
	avatar = &d.User.GetX(ctx, userID).QueryProfile().OnlyX(ctx).Avatar
	return
}

func (d *DB) GetLocalUserForForeignUser(userID uint64, host string) (localUserID uint64, err error) {
	defer doRecovery(&err)
	localUserID = d.ForeignUser.
		Query().
		Where(
			foreignuser.And(
				foreignuser.Foreignid(userID),
				foreignuser.Host(host),
			),
		).
		QueryUser().
		OnlyX(ctx).ID
	return
}

func (d *DB) getUserStem(user *entgen.User, profile *entgen.Profile) types.UserData {
	return types.UserData{
		UserID:   user.ID,
		Username: profile.Username,
		Avatar:   profile.Avatar,
		Status:   profile.Status,
		IsBot:    profile.IsBot,
	}
}

func (d *DB) GetUserByEmail(email string) (userData types.UserData, err error) {
	defer doRecovery(&err)
	localUser := d.LocalUser.Query().Where(localuser.Email(email)).WithUser().OnlyX(ctx)
	user := localUser.QueryUser().OnlyX(ctx)
	profile := user.QueryProfile().OnlyX(ctx)
	userData = d.getUserStem(user, profile)
	userData.Email = localUser.Email
	userData.Password = localUser.Password
	return
}

func (d *DB) GetUserByID(userID uint64) (userData types.UserData, err error) {
	defer doRecovery(&err)
	user := d.User.GetX(ctx, userID)
	profile := user.QueryProfile().OnlyX(ctx)
	userData = d.getUserStem(user, profile)
	return
}

func (d *DB) GetUsersByID(userID []uint64) (data []types.UserData, err error) {
	defer doRecovery(&err)

	dat := d.User.Query().Where(user.IDIn(userID...)).WithProfile().AllX(ctx)

	for _, it := range dat {
		data = append(data, d.getUserStem(it, it.Edges.Profile))
	}

	return
}

func (d *DB) GetUserMetadata(userID uint64, appID string) (meta string, err error) {
	defer doRecovery(&err)
	meta = d.User.
		Query().
		Where(
			user.ID(userID),
		).
		QueryMetadata().
		Where(
			usermeta.ID(appID),
		).
		OnlyX(ctx).
		Meta
	return
}

func (d *DB) updateProfileStem(userID uint64) *entgen.ProfileUpdate {
	return d.Profile.
		Update().
		Where(
			profile.HasUserWith(
				user.ID(userID)),
		)
}

func (d *DB) SetAvatar(userID uint64, avatar string) (err error) {
	defer doRecovery(&err)
	d.updateProfileStem(userID).SetAvatar(avatar)
	return
}

func (d *DB) SetIsBot(userID uint64, isBot bool) (err error) {
	defer doRecovery(&err)
	d.updateProfileStem(userID).
		SetIsBot(isBot)
	return
}

func (d *DB) SetStatus(userID uint64, status harmonytypesv1.UserStatus) (err error) {
	defer doRecovery(&err)
	d.updateProfileStem(userID).
		SetStatus(int16(status))
	return
}

func (d *DB) SetUsername(userID uint64, username string) (err error) {
	defer doRecovery(&err)
	d.updateProfileStem(userID).
		SetUsername(username)
	return
}

func (d *DB) UserIsLocal(userID uint64) (isLocal bool, err error) {
	defer doRecovery(&err)
	isLocal = d.User.Query().Where(user.ID(userID)).QueryLocalUser().ExistX(ctx)
	return
}

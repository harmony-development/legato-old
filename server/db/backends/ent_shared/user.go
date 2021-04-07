package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/foreignuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/localuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/harmony-development/legato/server/db/ent/entgen/usermeta"
	"github.com/harmony-development/legato/server/db/types"
)

func (d *database) AddForeignUser(host string, userID, localUserID uint64, username, avatar string) (err error) {
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

func (d *database) AddLocalUser(userID uint64, email, username string, passwordHash []byte) (err error) {
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

func (d *database) EmailExists(email string) (exists bool, err error) {
	defer doRecovery(&err)
	exists = d.LocalUser.Query().Where(localuser.Email(email)).ExistX(ctx)
	return
}

func (d *database) GetAvatar(userID uint64) (avatar *string, err error) {
	defer doRecovery(&err)
	avatar = &d.User.GetX(ctx, userID).QueryProfile().OnlyX(ctx).Avatar
	return
}

func (d *database) GetLocalUserForForeignUser(userID uint64, host string) (localUserID uint64, err error) {
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

func (d *database) getUserStem(user *entgen.User, profile *entgen.Profile) types.UserData {
	return types.UserData{
		UserID:   user.ID,
		Username: profile.Username,
		Avatar:   &profile.Avatar,
		Status:   profile.Status,
		IsBot:    profile.IsBot,
	}
}

func (d *database) GetUserByEmail(email string) (userData types.UserData, err error) {
	defer doRecovery(&err)
	localUser := d.LocalUser.Query().Where(localuser.Email(email)).WithUser().OnlyX(ctx)
	user := localUser.QueryUser().OnlyX(ctx)
	profile := user.QueryProfile().OnlyX(ctx)
	userData = d.getUserStem(user, profile)
	userData.Email = localUser.Email
	userData.Password = localUser.Password
	return
}

func (d *database) GetUserByID(userID uint64) (userData types.UserData, err error) {
	defer doRecovery(&err)
	user := d.User.GetX(ctx, userID)
	profile := user.QueryProfile().OnlyX(ctx)
	userData = d.getUserStem(user, profile)
	return
}

func (d *database) GetUserMetadata(userID uint64, appID string) (meta string, err error) {
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

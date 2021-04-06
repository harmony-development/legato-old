package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/foreignuser"
	"github.com/harmony-development/legato/server/db/ent/entgen/localuser"
)

func (d *database) AddForeignUser(host string, userID, localUserID uint64, username, avatar string) (err error) {
	doRecovery(&err)
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
	doRecovery(&err)
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
	doRecovery(&err)
	exists = d.LocalUser.Query().Where(localuser.Email(email)).ExistX(ctx)
	return
}

func (d *database) GetAvatar(userID uint64) (avatar *string, err error) {
	doRecovery(&err)
	avatar = &d.User.GetX(ctx, userID).QueryProfile().OnlyX(ctx).Avatar
	return
}

func (d *database) GetLocalUserForForeignUser(userID uint64, host string) (localUserID uint64, err error) {
	doRecovery(&err)
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

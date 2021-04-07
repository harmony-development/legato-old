package ent_shared

import (
	"time"

	"github.com/harmony-development/legato/server/db/ent/entgen/session"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

func (d *database) ExpireSessions() (err error) {
	defer doRecovery(&err)
	d.Session.Delete().Where(session.ExpiresLT(time.Now())).ExecX(ctx)
	return
}

func (d *database) ExtendSession(session string) (err error) {
	defer doRecovery(&err)
	d.Session.UpdateOneID(session).ExecX(ctx)
	return
}

func (d *database) AddSession(userID uint64, session string) (err error) {
	defer doRecovery(&err)

	d.Session.Create().
		SetUser(
			d.User.Query().
				Where(
					user.ID(userID),
					user.HasLocalUser(),
				).
				OnlyX(ctx),
		).
		SetID(session).
		SaveX(ctx)

	return
}

func (d *database) SessionToUserID(sid string) (userID uint64, err error) {
	defer doRecovery(&err)

	userID = d.Client.Session.
		Query().
		Where(session.ID(sid)).
		QueryUser().
		OnlyX(ctx).
		ID

	return
}

package sqlite

import "github.com/harmony-development/legato/server/db/backends/sqlite/ent/session"

func (d *database) AddSession(userID uint64, session string) (err error) {
	defer doRecovery(&err)

	user := d.User.GetX(ctx, userID).QueryLocalUser().OnlyX(ctx)
	d.Session.Create().SetUser(user).SetSessionID(session).SaveX(ctx)

	return
}

func (d *database) SessionToUserID(sid string) (uid uint64, err error) {
	defer doRecovery(&err)

	return uint64(d.Client.Session.Query().WithUser().Where(session.SessionID(sid)).OnlyX(ctx).Edges.User.ID), nil
}

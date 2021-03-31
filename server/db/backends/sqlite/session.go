//+build ignore
package sqlite

import (
	"github.com/harmony-development/legato/server/db/ent/entgen/session"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
	"github.com/ztrue/tracerr"
)

func (d *database) AddSession(userID uint64, session string) error {
	user, err := d.User.Query().Where(user.ID(userID), user.HasLocalUser()).Only(ctx)
	if err != nil {
		return tracerr.Wrap(err)
	}

	if _, err := d.Session.Create().SetUser(user).SetSessionid(session).Save(ctx); err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (d *database) SessionToUserID(sid string) (uint64, error) {
	s, err := d.Client.Session.
		Query().
		Where(session.Sessionid(sid)).
		QueryUser().
		Only(ctx)
	if err != nil {
		return 0, tracerr.Wrap(err)
	}
	return s.ID, nil
}

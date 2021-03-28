package sqlite

import (
	"database/sql"

	"github.com/harmony-development/legato/server/db/ent/entgen/localuser"
	"github.com/harmony-development/legato/server/db/queries"
)

func (d *database) EmailExists(email string) (bool, error) {
	return d.Client.LocalUser.Query().Where(localuser.Email(email)).Exist(ctx)
}

func (d *database) AddLocalUser(userID uint64, email, username string, passwordHash []byte) (rerr error) {
	defer doRecovery(&rerr)

	user := d.Client.User.Create().
		SetID(userID).
		SaveX(ctx)

	d.Client.Profile.Create().SetUser(user).SaveX(ctx)

	d.Client.LocalUser.
		Create().
		SetEmail(email).
		SetPassword(passwordHash).
		SetUser(user).
		SetUsername(username).
		SaveX(ctx)

	return
}

func (d *database) GetUserByEmail(email string) (q queries.GetUserByEmailRow, err error) {
	defer doRecovery(&err)

	user := d.Client.LocalUser.Query().Where(localuser.Email(email)).OnlyX(ctx)
	profile := user.QueryUser().QueryProfile().OnlyX(ctx)

	return queries.GetUserByEmailRow{
		UserID:   user.QueryUser().OnlyIDX(ctx),
		Email:    email,
		Username: user.Username,
		Avatar: sql.NullString{
			String: profile.Avatar,
		},
		Status:   profile.Status,
		Password: user.Password,
	}, nil
}

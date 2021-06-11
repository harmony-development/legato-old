package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/host"
)

func (d *DB) ensureHost(hs string) *entgen.Host {
	if d.Host.Query().Where(host.Host(hs)).ExistX(ctx) {
		return d.Host.Query().Where(host.Host(hs)).OnlyX(ctx)
	}
	return d.Host.Create().SetHost(hs).SaveX(ctx)
}
func (d *DB) GetHostQueue(host string) (data []byte, err error) {
	doRecovery(&err)

	return d.ensureHost(host).Eventqueue, nil
}
func (d *DB) SetHostQueue(host string, data []byte) (err error) {
	doRecovery(&err)

	d.ensureHost(host).Update().SetEventqueue(data).SaveX(ctx)

	return nil
}

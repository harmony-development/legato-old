package ent_shared

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/harmony-development/legato/server/db/types"
)

func TestGetPercentImplemented(t *testing.T) {
	db := &DB{}
	dbType := reflect.TypeOf(db)
	ifaceType := reflect.TypeOf((*types.IHarmonyDB)(nil)).Elem()
	meths := map[string]struct{}{}
	numImplemented := 0
	numNeeded := ifaceType.NumMethod()
	for i := 0; i < ifaceType.NumMethod(); i++ {
		meths[ifaceType.Method(i).Name] = struct{}{}
	}
	for i := 0; i < dbType.NumMethod(); i++ {
		if _, ok := meths[dbType.Method(i).Name]; ok {
			numImplemented++
		}
	}
	fmt.Printf("%f%% implemented\r\n", 100*(float32(numImplemented)/float32(numNeeded)))
}

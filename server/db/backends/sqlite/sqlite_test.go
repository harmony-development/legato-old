package sqlite

import (
	"testing"

	"github.com/harmony-development/legato/server/db/backends/ent_shared"
	"github.com/harmony-development/legato/server/db/ent/entgen"
)

// testing helpers
func createDatabase() *database {
	c, err := entgen.Open("sqlite3", "file:test.db?cache=shared&mode=memory&_fk=1")
	if err != nil {
		panic(err)
	}
	db, err := ent_shared.New(c, nil)
	if err != nil {
		panic(err)
	}
	return &database{
		DB: db,
	}
}

// makeLocalUser makes a local user for testing purposes
func makeLocalUser(db *database, userID uint64, email, username string, t *testing.T) {
	err := db.AddLocalUser(userID, email, username, []byte{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// TestEmotePackOperations tests various operations on emote packs
func TestEmotePackOperations(t *testing.T) {
	db := createDatabase()
	defer db.DB.Close()

	userID := uint64(1)
	packID := uint64(2)
	otherPackID := uint64(3)
	email := "aa@aa.aa"
	username := "aa"

	makeLocalUser(db, userID, email, username, t)

	// create a new pack
	err := db.CreateEmotePack(userID, packID, "testing")
	if err != nil {
		t.Error(err)
	}

	// try to create a pack with the same name
	err = db.CreateEmotePack(userID, otherPackID, "testing")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// create a new pack with a different name
	err = db.CreateEmotePack(userID, otherPackID, "testing2")
	if err != nil {
		t.Error(err)
	}

	// add an emote to the pack
	err = db.AddEmoteToPack(packID, "hmc://foo", "emote1")
	if err != nil {
		t.Error(err)
	}

	// try to add an emote to the pack again
	err = db.AddEmoteToPack(packID, "hmc://idk", "emote1")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// add another emote to the pack
	err = db.AddEmoteToPack(packID, "hmc://foo", "emote2")
	if err != nil {
		t.Error(err)
	}

	// delete the emote from the pack
	err = db.DeleteEmoteFromPack(packID, "emote1")
	if err != nil {
		t.Error(err)
	}

	// delete the pack
	err = db.DeleteEmotePack(packID)
	if err != nil {
		t.Error(err)
	}

	// make sure emote2 & the pack is deleted
	err = db.DeleteEmoteFromPack(packID, "emote2")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

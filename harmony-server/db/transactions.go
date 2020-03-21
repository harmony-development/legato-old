package db

import (
	"github.com/thanhpk/randstr"
)

func DeleteChannelTransaction(guildid string, channelid string) error {
	transaction, err := DBInst.Begin()
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM messages WHERE channelid=$1 AND guildid=$2", channelid, guildid)
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM channels WHERE channelid=$1 AND guildid=$2", channelid, guildid)
	if err != nil {
		return err
	}
	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func CreateGuildTransaction(guildname string, owner string) (*string, error) {
	guildid := randstr.Hex(16)
	createGuildTransaction, err := DBInst.Begin()
	if err != nil {
		return nil, err
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO guilds(guildid, guildname, picture, owner) VALUES($1, $2, $3, $4);`, guildid, guildname, "", owner)
	if err != nil {
		return nil, err
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO guildmembers(userid, guildid) VALUES($1, $2);`, owner, guildid)
	if err != nil {
		return nil, err
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO channels(channelid, guildid, channelname) VALUES($1, $2, $3)`, randstr.Hex(16), guildid, "general")
	if err != nil {
		return nil, err
	}
	err = createGuildTransaction.Commit()
	if err != nil {
		return nil, err
	}
	return &guildid, nil
}

func DeleteGuildTransaction(guildid string) error {
	transaction, err := DBInst.Begin()
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM messages WHERE guildid=$1", guildid)
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM channels WHERE guildid=$1", guildid)
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM  invites WHERE guildid=$1", guildid)
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM guildmembers WHERE guildid=$1", guildid)
	if err != nil {
		return err
	}
	_, err = transaction.Exec("DELETE FROM guilds WHERE guildid=$1", guildid)
	if err != nil {
		return err
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func DeleteMessageTransaction(guild string, channel string, message string, userid string) error {
	transaction, err := DBInst.Begin()
	if err != nil {
		return err
	}
	res, err := transaction.Exec("DELETE FROM messages WHERE guildid=$1 AND channelid=$2 AND messageid=$3 AND (author=$4 OR (SELECT owner FROM guilds WHERE guildid=$5)=$6)", guild, channel, message, userid, guild, userid)
	if err != nil {
		return err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCount != 1 {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return nil
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func JoinGuildTransaction(userid string, guildid string, inviteCode string) error {
	joinGuildTransaction, err := DBInst.Begin()
	if err != nil {
		return err
	}
	_, err = joinGuildTransaction.Exec("INSERT INTO guildmembers(userid, guildid) VALUES($1, $2)", userid, guildid)
	if err != nil {
		return err
	}
	_, err = joinGuildTransaction.Exec("UPDATE invites SET invitecount=invitecount+1 WHERE inviteid=$1", inviteCode)
	if err != nil {
		return err
	}
	err = joinGuildTransaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func IsUserOwner(guildid string, userid string) bool {
	err := DBInst.QueryRow("SELECT guilds.owner FROM guilds WHERE guildid=$1 AND guilds.owner=$2", guildid, userid)
	return err == nil
}
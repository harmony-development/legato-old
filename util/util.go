package util

import (
	"database/sql"
	"strconv"
	"time"
)

func U64TSA(uint64s []uint64) (ret []string) {
	for _, number := range uint64s {
		ret = append(ret, strconv.FormatUint(number, 10))
	}
	return
}

func U64TS(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func TimeTS(tiempo time.Time) string {
	return strconv.FormatInt(tiempo.UTC().Unix(), 10)
}

func NullTimeTS(tiempo sql.NullTime) (ret *string) {
	if tiempo.Valid {
		formatted := TimeTS(tiempo.Time)
		ret = &formatted
	}
	return
}

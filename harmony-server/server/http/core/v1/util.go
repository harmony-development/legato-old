package v1

import (
	"database/sql"
	"strconv"
	"time"
)

func u64TSA(uint64s []uint64) (ret []string) {
	for _, number := range uint64s {
		ret = append(ret, strconv.FormatUint(number, 10))
	}
	return
}

func u64TS(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func timeTS(tiempo time.Time) string {
	return strconv.FormatInt(tiempo.UTC().Unix(), 10)
}

func nullTimeTS(tiempo sql.NullTime) (ret *string) {
	if tiempo.Valid {
		formatted := timeTS(tiempo.Time)
		ret = &formatted
	}
	return
}

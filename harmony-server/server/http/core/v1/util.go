package v1

import "strconv"

func u64TSA(uint64s []uint64) (ret []string) {
	for _, number := range uint64s {
		ret = append(ret, strconv.FormatUint(number, 10))
	}
	return
}

func u64TS(number uint64) string {
	return strconv.FormatUint(number, 10)
}

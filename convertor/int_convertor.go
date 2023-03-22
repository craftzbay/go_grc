package convertor

import "strconv"

func UintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

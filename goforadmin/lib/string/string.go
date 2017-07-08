package string

import (
	"strconv"
	"strings"
)

func SplitToInt64Slice(s string, sep string) []int64 {
	var ret []int64

	ss := strings.Split(s, sep)
	if len(ss) > 0 {
		ret = make([]int64, 0, len(ss))
	}

	for _, v := range ss {
		i, _ := strconv.ParseInt(v, 10, 64)
		ret = append(ret, i)
	}

	return ret
}

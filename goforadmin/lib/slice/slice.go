package slice

func InIntArray(val int, arr []int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}

	return false
}

func InInt64Array(val int64, arr []int64) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}

	return false
}

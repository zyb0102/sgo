package utils

// IsValInList 判断一个list里是否存在某值
func IsValInList(v string, list []string) bool {
	for _, val := range list {
		if v == val {
			return true
		}
	}
	return false
}

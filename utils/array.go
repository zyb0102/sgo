package utils

func IsValInList(v string, list []string) bool {
	for _, val := range list {
		if v == val {
			return true
		}
	}
	return false
}

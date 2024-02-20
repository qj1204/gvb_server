package utils

// InList 判断字符串是否在列表中
func InList(key string, list []string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

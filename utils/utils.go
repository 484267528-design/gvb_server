package utils

func InList(key string, list []string) bool {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}

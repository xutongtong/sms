package utils

func InArray(key string, params []interface{}) bool {
	for _, v := range params {
		if key == v.(string) {
			return true
		}
	}

	return false
}
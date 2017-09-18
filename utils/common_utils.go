package utils

func UnpackString(list []interface{}) []interface{} {
	vals := make([]interface{}, len(list))
	for k, v := range list {
		vals[k] = v
	}
	return vals
}

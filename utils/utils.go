package utils

import "fmt"

func GetIntField(hit map[string]interface{}, field string) int {
	val, ok := hit[field]
	if !ok {
		return 0
	}
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case int64:
		return int(v)
	default:
		panic(fmt.Sprintf("unexpected type %T for %s field", v, field))
	}
}

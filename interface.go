package helper

import "fmt"

// TypeSwitch 类型判断
func TypeSwitch(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%T", v)
	}
}

// CopyMap 深拷贝 map
func CopyMap[M ~map[K]V, K comparable, V any](m M) M {
	result := make(M, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

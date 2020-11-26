package util

import (
	"fmt"
	"strconv"
)

// InterfaceToString 接口转其他类型
func InterfaceToString(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int:
		return strconv.Itoa(inter.(int))
	case int64:
		return strconv.FormatInt(inter.(int64), 64)
	case float64:
		return strconv.FormatFloat(inter.(float64), byte('f'), 4, 64)
	default :
		return fmt.Sprintf("%v", inter)
	}
}
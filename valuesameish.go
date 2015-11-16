package valuesameish

import (
	"fmt"
	"strconv"
)

func Sameish(first interface{}, second interface{}) bool {
	switch first.(type) {
	case int:
		switch second.(type) {
		case int:
			return first.(int) == second.(int)
		case int8:
			return first.(int8) == second.(int8)
		case int16:
			return first.(int16) == second.(int16)
		case int32:
			return first.(int32) == second.(int32)
		case int64:
			return first.(int64) == second.(int64)
		case uint8:
			return first.(uint8) == second.(uint8)
		case uint16:
			return first.(uint16) == second.(uint16)
		case uint32:
			return first.(uint32) == second.(uint32)
		case uint64:
			return first.(uint64) == second.(uint64)
		case float32:
			return first.(int) == int(second.(float32))
		case float64:
			return first.(int) == int(second.(float64))
		case string:
			//Maybe the thing in string is an int
			i, _ := strconv.Atoi(second.(string))
			if first.(int) == i {
				return true
			}
			//Maybe the thing in string is a float
			f, _ := strconv.ParseFloat(second.(string), 64)
			return first.(int) == int(f)
		default:
			return false
		}
	case string:
		switch second.(type) {
		case string:
			return first.(string) == second.(string)
		case int:
			return Sameish(second, first)
		case float32:
			return Sameish(second, first)
		case float64:
			return Sameish(second, first)
		default:
			return false
		}
	default:
		return false
	}
	fmt.Println("Oh yeah")
	return false
}

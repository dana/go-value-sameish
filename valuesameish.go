package valuesameish

import (
	"fmt"
	"strconv"
)

func sameishInt64String(first int64, second string) bool {
	//Maybe the thing in string is an int
	i, _ := strconv.Atoi(second)
	if first == int64(i) {
		return true
	}
	//Maybe the thing in string is a float
	f, _ := strconv.ParseFloat(second, 64)
	return first == int64(f)
}

func Sameish(first interface{}, second interface{}) bool {
	switch first.(type) {
	case int:
		switch second.(type) {
		case int:
			return int64(first.(int)) == int64(second.(int))
		case int8:
			return int64(first.(int)) == int64(second.(int8))
		case int16:
			return int64(first.(int)) == int64(second.(int16))
		case int32:
			return int64(first.(int)) == int64(second.(int32))
		case int64:
			return int64(first.(int)) == second.(int64)
		case uint8:
			return uint64(first.(int)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(int)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(int)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(int)) == second.(uint64)
		case float32:
			return int64(first.(int)) == int64(second.(float32))
		case float64:
			return int64(first.(int)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(int)), second.(string))
		default:
			return false
		}
	case int8:
		switch second.(type) {
		case int:
			return int64(first.(int8)) == int64(second.(int))
		case int8:
			return int64(first.(int8)) == int64(second.(int8))
		case int16:
			return int64(first.(int8)) == int64(second.(int16))
		case int32:
			return int64(first.(int8)) == int64(second.(int32))
		case int64:
			return int64(first.(int8)) == second.(int64)
		case uint8:
			return uint64(first.(int8)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(int8)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(int8)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(int8)) == second.(uint64)
		case float32:
			return int64(first.(int8)) == int64(second.(float32))
		case float64:
			return int64(first.(int8)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(int8)), second.(string))
		default:
			return false
		}
	case int16:
		switch second.(type) {
		case int:
			return int64(first.(int16)) == int64(second.(int))
		case int8:
			return int64(first.(int16)) == int64(second.(int8))
		case int16:
			return int64(first.(int16)) == int64(second.(int16))
		case int32:
			return int64(first.(int16)) == int64(second.(int32))
		case int64:
			return int64(first.(int16)) == second.(int64)
		case uint8:
			return uint64(first.(int16)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(int16)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(int16)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(int16)) == second.(uint64)
		case float32:
			return int64(first.(int16)) == int64(second.(float32))
		case float64:
			return int64(first.(int16)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(int16)), second.(string))
		default:
			return false
		}
	case int32:
		switch second.(type) {
		case int:
			return int64(first.(int32)) == int64(second.(int))
		case int8:
			return int64(first.(int32)) == int64(second.(int8))
		case int16:
			return int64(first.(int32)) == int64(second.(int16))
		case int32:
			return int64(first.(int32)) == int64(second.(int32))
		case int64:
			return int64(first.(int32)) == second.(int64)
		case uint8:
			return uint64(first.(int32)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(int32)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(int32)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(int32)) == second.(uint64)
		case float32:
			return int64(first.(int32)) == int64(second.(float32))
		case float64:
			return int64(first.(int32)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(int32)), second.(string))
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

/*
matrix of types

int int
int float
int string

*float int
float float
float string

*string int
*string float
string string

*/

package valuesameish

import (
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

//a lot of the below code has been mechanically created and I am sure
//that a lot of it is incorrect outside of 'middle' cases
//TODO: figure out and then encode what 'Sameish' means for the full matric
//of types
func Sameish(first interface{}, second interface{}) bool {
	switch first.(type) {
	case uint:
		switch second.(type) {
		case int:
			return int64(first.(uint)) == int64(second.(int))
		case int8:
			return int64(first.(uint)) == int64(second.(int8))
		case int16:
			return int64(first.(uint)) == int64(second.(int16))
		case int32:
			return int64(first.(uint)) == int64(second.(int32))
		case int64:
			return int64(first.(uint)) == second.(int64)
		case uint8:
			return uint64(first.(uint)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(uint)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(uint)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(uint)) == second.(uint64)
		case float32:
			return int64(first.(uint)) == int64(second.(float32))
		case float64:
			return int64(first.(uint)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(uint)), second.(string))
		default:
			return false
		}
	case uint8:
		switch second.(type) {
		case int:
			return int64(first.(uint8)) == int64(second.(int))
		case int8:
			return int64(first.(uint8)) == int64(second.(int8))
		case int16:
			return int64(first.(uint8)) == int64(second.(int16))
		case int32:
			return int64(first.(uint8)) == int64(second.(int32))
		case int64:
			return int64(first.(uint8)) == second.(int64)
		case uint8:
			return uint64(first.(uint8)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(uint8)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(uint8)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(uint8)) == second.(uint64)
		case float32:
			return int64(first.(uint8)) == int64(second.(float32))
		case float64:
			return int64(first.(uint8)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(uint8)), second.(string))
		default:
			return false
		}
	case uint16:
		switch second.(type) {
		case int:
			return int64(first.(uint16)) == int64(second.(int))
		case int8:
			return int64(first.(uint16)) == int64(second.(int8))
		case int16:
			return int64(first.(uint16)) == int64(second.(int16))
		case int32:
			return int64(first.(uint16)) == int64(second.(int32))
		case int64:
			return int64(first.(uint16)) == second.(int64)
		case uint8:
			return uint64(first.(uint16)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(uint16)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(uint16)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(uint16)) == second.(uint64)
		case float32:
			return int64(first.(uint16)) == int64(second.(float32))
		case float64:
			return int64(first.(uint16)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(uint16)), second.(string))
		default:
			return false
		}
	case uint32:
		switch second.(type) {
		case int:
			return int64(first.(uint32)) == int64(second.(int))
		case int8:
			return int64(first.(uint32)) == int64(second.(int8))
		case int16:
			return int64(first.(uint32)) == int64(second.(int16))
		case int32:
			return int64(first.(uint32)) == int64(second.(int32))
		case int64:
			return int64(first.(uint32)) == second.(int64)
		case uint8:
			return uint64(first.(uint32)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(uint32)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(uint32)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(uint32)) == second.(uint64)
		case float32:
			return int64(first.(uint32)) == int64(second.(float32))
		case float64:
			return int64(first.(uint32)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(uint32)), second.(string))
		default:
			return false
		}
	case uint64:
		switch second.(type) {
		case int:
			return int64(first.(uint64)) == int64(second.(int))
		case int8:
			return int64(first.(uint64)) == int64(second.(int8))
		case int16:
			return int64(first.(uint64)) == int64(second.(int16))
		case int32:
			return int64(first.(uint64)) == int64(second.(int32))
		case int64:
			return int64(first.(uint64)) == second.(int64)
		case uint8:
			return uint64(first.(uint64)) == uint64(second.(uint8))
		case uint16:
			return uint64(first.(uint64)) == uint64(second.(uint16))
		case uint32:
			return uint64(first.(uint64)) == uint64(second.(uint32))
		case uint64:
			return uint64(first.(uint64)) == second.(uint64)
		case float32:
			return int64(first.(uint64)) == int64(second.(float32))
		case float64:
			return int64(first.(uint64)) == int64(second.(float64))
		case string:
			return sameishInt64String(int64(first.(uint64)), second.(string))
		default:
			return false
		}
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
		case int16:
			return Sameish(second, first)
		case int32:
			return Sameish(second, first)
		case int64:
			return Sameish(second, first)
		case uint:
			return Sameish(second, first)
		case uint8:
			return Sameish(second, first)
		case uint16:
			return Sameish(second, first)
		case uint32:
			return Sameish(second, first)
		case uint64:
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

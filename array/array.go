package array

import (
	"fmt"
	"reflect"
)

// Find index of value in array
// TypeError return -2
// NotFind return -1
func ArrayIndexOf(num, array interface{}) int {
	va := reflect.ValueOf(array)
	vaType := reflect.TypeOf(array).Kind()
	switch vaType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < va.Len(); i++ {
			if num == va.Index(i).Interface() {
				return i
			}
		}
		return -1
	default:
		return -2
	}
}

// DeleteDuplicate delete duplicate value in array
func DeleteDuplicate(a interface{}) (ret []interface{}, err error) {
	if a == nil {
		return nil, fmt.Errorf("nil input")
	}
	va := reflect.ValueOf(a)
	vaType := reflect.TypeOf(a).Kind()
	switch vaType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < va.Len(); i++ {
			res := ArrayIndexOf(va.Index(i).Interface(), ret)
			if res == -1 {
				ret = append(ret, va.Index(i).Interface())
			}
			if res == -2 {
				return nil, fmt.Errorf("type error")
			}
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("type error, please input array")
	}
}

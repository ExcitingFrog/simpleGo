package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// unsafe
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// unsafe
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// ConvertInterfaceArrayToFloat64Array ...
func ConvertInterfaceArrayToFloat64Array(array []interface{}) ([]float64, error) {
	if len(array) == 0 {
		return nil, fmt.Errorf("array length 0")
	}
	res := []float64{}
	for _, value := range array {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Float64:
			res = append(res, value.(float64))
		case reflect.Int:
			res = append(res, float64(value.(int)))
		default:
			return nil, fmt.Errorf("type error")
		}
	}
	return res, nil
}

// ConvertStringArrayToInt64Array ...
func ConvertStringArrayToInt64Array(array []string) (res []int64, err error) {
	for _, value := range array {
		num, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		res = append(res, num)
	}
	return res, nil
}

// ConvertInterfaceArrayToInt64Array ...
func ConvertInterfaceArrayToInt64Array(array []interface{}) ([]int64, error) {
	if len(array) == 0 {
		return nil, fmt.Errorf("array length 0")
	}
	res := []int64{}
	for _, value := range array {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Int64:
			res = append(res, value.(int64))
		case reflect.Float64:
			res = append(res, int64(value.(float64)))
		case reflect.Int:
			res = append(res, int64(value.(int)))
		default:
			return nil, fmt.Errorf("type error")
		}
	}
	return res, nil
}

// ConvertInterfaceArrayToStringArray ...
func ConvertInterfaceArrayToStringArray(array []interface{}) ([]string, error) {
	if len(array) == 0 {
		return nil, fmt.Errorf("array length 0")
	}
	res := []string{}
	for _, value := range array {
		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			res = append(res, value.(string))
		default:
			return nil, fmt.Errorf("type error")
		}
	}
	return res, nil
}

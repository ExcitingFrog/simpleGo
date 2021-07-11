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
		case reflect.String:
			num, err := strconv.ParseFloat(value.(string), 64)
			if err != nil {
				return nil, err
			}
			res = append(res, num)
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
		case reflect.String:
			num, err := strconv.ParseFloat(value.(string), 64)
			if err != nil {
				return nil, err
			}
			res = append(res, int64(num))
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
		case reflect.Float64:
			res = append(res, fmt.Sprintf("%g", value.(float64)))
		case reflect.Int:
			res = append(res, strconv.Itoa(value.(int)))
		case reflect.Int64:
			res = append(res, strconv.FormatInt(value.(int64), 64))
		default:
			return nil, fmt.Errorf("type error")
		}
	}
	return res, nil
}

// Decimal ...
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// DecimalToZero ...
func DecimalToZero(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.0f", value), 64)
	return value
}

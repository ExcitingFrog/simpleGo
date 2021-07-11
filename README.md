## SimpleGo

This repository contains some functions in common use.It can reduce time wasted on some esay question or annoying type problem.

### How to use

```go
go get -u github.com/ExcitingFrog/simpleGo
```

### convert

#### StringToBytes

Convert string to byte quick but unsafe

```go
convert.StringToBytes("hello,go")
// [104 101 108 108 111 44 103 111]
```

#### BytesToString

Convert byte to string quick but unsafe

```go
convert.BytesToString([]byte{104, 101, 108, 108, 111, 44, 103, 111})
// hello,go
```

#### ConvertStringArrayToInt64Array

Convert string array to int64 array

```go
convert.ConvertStringArrayToInt64Array([]string{"1", "2", "3"})
// [1 2 3] <nil>
```

#### ConvertInterfaceArrayToFloat64Array

Convert interface array to float64 array

```go
convert.ConvertInterfaceArrayToFloat64Array([]interface{}{"1", 1.2, 3, "2.2"})
// [1 1.2 3 2.2] <nil>
```

#### ConvertInterfaceArrayToInt64Array

Convert interface array to int64 array

```go
convert.ConvertInterfaceArrayToInt64Array([]interface{}{3, 1.2, "1.1", "4"})
// [3 1 1 4] <nil>
```

#### ConvertInterfaceArrayToStringArray

Convert interface array to string array

```go
convert.ConvertInterfaceArrayToStringArray([]interface{}{"he", 2, 1.2, "4.4"})
// [he 2 1.2 4.4] <nil>
```

#### Decimal

Convert float64 to two decimal places

```go
fmt.Println(convert.Decimal(1.3243344))
// 1.32
```

#### DecimalToZero

Convert float64 to two decimal places

```go
fmt.Println(convert.DecimalToZero(1.7))
// 2
```



### random

#### RandStringRunes

Generate assigned length random array its content comes from assigned rune

```go
fmt.Println(random.RandStringRunes(10, random.NumberRunes))
// 7907883425
fmt.Println(random.RandStringRunes(10, random.LetterNumberRunes))
// G0WhnjY3js
```



 ### array

#### ArrayIndexOf

Get index of value in one array.

Type error return -2

Not Find return -1

```go
fmt.Println(array.ArrayIndexOf(1, []int{2, 4, 1}))
// 2
fmt.Println(array.ArrayIndexOf(1.32, []float64{2.31, 4, 1}))
// -1
fmt.Println(array.ArrayIndexOf(1.32, []float64{2.31, 1.32, 1}))
// 1
fmt.Println(array.ArrayIndexOf("1.32", 1.32))
// -2
```

#### DeleteDuplicate

Delete duplicate value in one array.It can also used in pointer struct.

```go
fmt.Println(array.DeleteDuplicate([]int64{1, 1, 3, 1}))
// [1 3] <nil>
fmt.Println(array.DeleteDuplicate([]string{"1", "2", "3", "3"}))
// [1 2 3] <nil>
```



### stime

#### WeekStart

Get unix time that first Monday start in one year.

```go
fmt.Println(stime.WeekStart(2021, 1))
// 2021-01-04 00:00:00 +0800 CST
```

####  ParseDate

Parse string time to unix time.

```go
fmt.Println(stime.ParseDate("2006-01-02", "2021-01-04", "2021-01-11"))
// 1609689600 1610380799 <nil>
```


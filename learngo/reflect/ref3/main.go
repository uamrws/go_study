package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int
	fmt.Println(a)
	fmt.Println(reflect.ValueOf(a).IsValid())
	fmt.Println(reflect.ValueOf(a).IsNil())
	fmt.Println(reflect.ValueOf(a).IsZero())
	var b []int
	fmt.Println(b)
	fmt.Println(reflect.ValueOf(b).IsValid())
	fmt.Println(reflect.ValueOf(b).IsNil())
	fmt.Println(reflect.ValueOf(b).IsZero())
	var c struct{}
	fmt.Println("c:\n", c)
	fmt.Println(reflect.ValueOf(c).FieldByName("haha").IsValid())
	fmt.Println(reflect.ValueOf(c).MethodByName("haha").IsValid())
	//fmt.Println(reflect.ValueOf(c).IsNil())
	fmt.Println(reflect.ValueOf(c).IsZero())
	var d map[string]int
	fmt.Println("d:\n", d)
	fmt.Println(reflect.ValueOf(d).MapIndex(reflect.ValueOf("haha")).IsValid())
}

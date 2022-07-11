package main

import (
	"fmt"
	"reflect"
)

func modifyValue(a interface{}) {
	v := reflect.ValueOf(a)
	kind := v.Kind()
	fmt.Printf("%T\n", v)
	fmt.Println(kind)
	if kind == reflect.Ptr {
		v.Elem().SetInt(13)
	}
}

func main() {
	var a int64 = 100
	modifyValue(&a)
	fmt.Println(a)
}

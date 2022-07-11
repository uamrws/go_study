package main

import (
	"fmt"
	"reflect"
)

type some struct {
	name string
	age  uint8
}

func reflectTypeof(someOne interface{}) interface{} {
	t := reflect.TypeOf(someOne)
	fmt.Printf("type:%s, name:%s, king:%s, value:%#v\n", t, t.Name(), t.Kind(), reflect.ValueOf(someOne))
	return reflect.ValueOf(someOne)
}
func main() {
	var a int64 = 12
	var b int32 = 12
	var c int8 = 12
	d := map[string]int{"as": 1}
	e := []int{1, 2, 3}
	f := [3]int{1, 2, 3}
	g := some{"haha", 12}
	reflectTypeof(a)
	reflectTypeof(b)
	reflectTypeof(c)
	reflectTypeof(d)
	reflectTypeof(e)
	reflectTypeof(f)
	reflectTypeof(&f)
	reflectTypeof(reflectTypeof(g))
}

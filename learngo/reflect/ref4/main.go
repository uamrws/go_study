package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `haha:"name" shuaige:"nb"`
	Age  int
}

func (p *person) Eat(food string) {
	fmt.Printf("person %s eat %s\n", p.Name, food)
}
func main() {
	p := &person{"abc", 24}
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)
	fmt.Printf("name:%s, kind:%s\n", pt.Name(), pt.Kind())
	for i := 0; i < pt.Elem().NumField(); i++ {
		field := pt.Elem().Field(i)
		fmt.Printf(
			"name:%s, index:%d, type:%s, tag:%s\n",
			field.Name, field.Index, field.Type, field.Tag.Get("shuaige"),
		)
		if i == 0 {
			pv.Elem().Field(i).SetString("asd")
		}
	}
	fmt.Println(*p)
	fmt.Println(pt.NumMethod())
	for i := 0; i < pt.NumMethod(); i++ {
		methodType := pt.Method(i)
		method := pv.Method(i)
		fmt.Printf(
			"name:%s, index:%d, type:%s, func:%s\n",
			methodType.Name, methodType.Index, methodType.Type, method.Type(),
		)
		methodType.Func.Call([]reflect.Value{pv, reflect.ValueOf("hh")})
		method.Call([]reflect.Value{reflect.ValueOf("jj")})
	}
}

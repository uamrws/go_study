package main

import (
	"fmt"
	"go_study/temp/t1"
	_ "go_study/temp/t1"
	_ "unsafe"
)

type B struct {
	t1.C
}

func (b B) H() {
	fmt.Println("B, H call")
}

func (b B) h() {
	fmt.Println("B, h call")
}

//调用其他包的私有方法或属性
//go:linkname cc go_study/temp/t1.cc
func cc()

func main() {
	// 这个是正确的
	b := B{}
	t1.N(b)
	cc()
}

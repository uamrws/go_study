package t1

import "fmt"

type A interface {
	H()
	h()
}

type C struct {
}

func (c C) H() {
	fmt.Println("A, H call")
}
func (c C) h() {
	fmt.Println("A, h call")
}

func N(a A) {
	a.h()
	a.H()
}

func cc() {
	fmt.Println("hahahah")
}

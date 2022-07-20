package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

func covertToBin(n int) string {
	s := ""
	for ; n > 0; n /= 2 {
		s = strconv.Itoa(n%2) + s
	}
	return s
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
func apply(op func(a, b int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
func modify(cache Cache) {
	cache.a = 2
}

type Cache struct {
	a int
}

func updateArr(arr []int) {
	arr[0] = 100
}
func main() {
	//fmt.Println(
	//	covertToBin(100),
	//)
	//printFile("temp.txt")
	//apply(func(a, b int) int {
	//	return int(math.Pow(float64(a), float64(b)))
	//}, 3, 4)
	//
	//a, b := 3, 4
	//swap(&a, &b)
	//fmt.Println(a, b)
	//c := 1
	//cache := Cache{a: c}
	//modify(cache)
	//fmt.Println(cache.a)
	//
	//m := make(map[int]int)
	//fmt.Println(m)
	//var aaa [5]int
	//fmt.Printf("%T %v", aaa, aaa)
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//s3 := append(s2, 10)
	//s4 := append(s3, 11)
	//s5 := append(s4, 12)
	//
	//fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))
	//fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	//fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	//fmt.Printf("s1,s2,s3,s4,s5,arr= %v,%v,%v,%v,%v,%v", s1, s2, s3, s4, s5, arr)

}

package main

import (
	"fmt"
	"go_study/learngo/retriver/mock"
	real2 "go_study/learngo/retriver/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"hahaha"}
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}

	fmt.Println(r.(*real2.Retriever).UserAgent)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println(v.Contents)
	default:
		fmt.Println(v.(*real2.Retriever).UserAgent)
	}
}

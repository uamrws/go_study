package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/info", info)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatalln("failed to listen 8088")
	}

}

func info(writer http.ResponseWriter, request *http.Request) {
	//tem, err := template.ParseFiles("learngo/http_template/template/info.html")
	//tem.Funcs(template.FuncMap{"hello": func(item string) (string, error) {
	//	return "hello " + item, nil
	//}})
	tem, err := template.New("info.html").Funcs(template.FuncMap{"hello": func(item string) (string, error) {
		return "hello " + item, nil
	}}).ParseFiles("learngo/http_template/template/info.html")
	if err != nil {
		log.Println("failed to open template file")
	}
	user := []struct {
		Name  string
		Phone string
	}{
		{"你好", "12341a2314"},
		{"你好2", "12341a2314"},
		{"你好3", "12341a2314"},
		{"你好4", "12341a2314"},
	}
	//if err = tem.Execute(writer, `<script>alert("你好我的主人")</script>`); err != nil {
	//	log.Println("处理异常")
	//}
	if err = tem.Execute(writer, user); err != nil {
		log.Printf("处理异常%v", err)
	}
}

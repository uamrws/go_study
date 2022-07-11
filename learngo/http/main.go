package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Mobile Safari/537.36",
	)
	//resp, err := http.DefaultClient.Do(request)
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("redirected:", req)
		fmt.Println(via)
		return nil
	}}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", s)

}

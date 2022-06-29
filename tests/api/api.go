package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func testPost() {
	urlStr := "https://tool.runoob.com/compile2.php"
	values := url.Values{}
	values.Add("code", "print('hello world')")
	values.Add("token", "4381fe197827ec87cbac9552f14ec62a")
	values.Add("stdin", "")
	values.Add("language", "15")
	values.Add("fileext", "py3")
	//建立请求链接
	r, _ := http.PostForm(urlStr, values)
	defer r.Body.Close()
	//读取buffer
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
}

func main() {
	testPost()
}

package service

import (
	"codecloud/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiService struct {
}

func (receiver *ApiService) Python(code string, response domain.Res) domain.Res {
	urlStr := "https://tool.runoob.com/compile2.php"
	values := url.Values{}
	values.Add("code", code)
	values.Add("token", "4381fe197827ec87cbac9552f14ec62a")
	values.Add("stdin", "")
	values.Add("language", "15")
	values.Add("fileext", "py3")
	//建立请求链接
	r, _ := http.PostForm(urlStr, values)
	defer r.Body.Close()
	//读取buffer
	Json, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(Json), &response)
	if err != nil {
		fmt.Println("json解析失败")
	}
	return response
}

func (receiver *ApiService) Golang(code string, response domain.Res) domain.Res {
	urlStr := "https://tool.runoob.com/compile2.php"
	values := url.Values{}
	values.Add("code", code)
	values.Add("token", "4381fe197827ec87cbac9552f14ec62a")
	values.Add("stdin", "")
	values.Add("language", "6")
	values.Add("fileext", "go")
	//建立请求链接
	r, _ := http.PostForm(urlStr, values)
	defer r.Body.Close()
	//读取buffer
	Json, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(Json), &response)
	if err != nil {
		fmt.Println("json解析失败")
	}
	return response
}

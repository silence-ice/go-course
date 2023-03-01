package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

func main() {
	BingFaRequse()
	//GetRequest()
	//PostRequest()
}

func PostRequest() {
	//声明传参结构
	type PostRquest struct {
		WechatAppId string `json:"wechatAppId"`
		Page        string `json:"page"`
	}

	data, _ := json.Marshal(&PostRquest{
		WechatAppId: "wx638f492187c8bfb2",
		Page:        "1",
	})

	log.Println(string(data))

	//form 表单数据
	contentType := "application/x-www-form-urlencoded"
	postData := "wechatAppId=wx638f492187c8bfb2&age=18"

	// json 数据
	//contentType := "application/json"
	//postData := string(data)
	resp, err := http.Post("https://api.banchengyun.com/mall/goods/discoverEntrance", contentType, strings.NewReader(postData))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}

	fmt.Println(string(b))
}

func GetRequest() {
	//请求地址
	requestHost := "https://api.banchengyun.com/mall/goods/discoverEntrance"

	u, err := url.ParseRequestURI(requestHost)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}

	//log.Println(u)

	//组合url数据
	data := url.Values{}
	data.Set("wechatAppId", "wx638f492187c8bfb2")
	u.RawQuery = data.Encode() // URL encode
	resp, err := http.Get(u.String())

	//Get请求
	//resp, err := http.Get(requestHost+"/mall/goods/discoverEntrance?wechatAppId=wx638f492187c8bfb2")
	//fmt.Println(resp,err)
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Print(string(body))
}

func BingFaWrite(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	GetRequest()
	wg.Done()
}

func BingFaRequse() {
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)

	for i := 0; i < num; i++ {
		go BingFaWrite(strconv.Itoa(i), &wg)

		//go func(i int) {
		//	fmt.Println(i)
		//	wg.Done()
		//}(i)
	}
	wg.Wait()
}

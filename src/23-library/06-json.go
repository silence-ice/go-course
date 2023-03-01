package main

import (
	"encoding/json"
	"fmt"
	//jsoniter "github.com/json-iterator/go"
)

/*
*

	1、声明json的结构
	2、tag就是标签，给结构体的每个字段打上一个标签，标签冒号前是类型，后面是标签名
*/
type User struct {
	Name    string `json:"name"` //将key的Name转化成name
	age     int    //key首字母没有大写，所以不能被访问，json_encode是不会出来的
	Id      int    `json:"id,string"`         //将key的Id转化成id，同时类型转成string
	Address string `json:"address,omitempty"` //omitempty 可以在序列化的时候忽略0值或者空值，但是要跟在key的后面
	Family  string `json:"-"`                 //-表示不进行json_encode
}

func main() {
	//json_encode
	u := User{
		Name:    "huangbin",
		age:     12,
		Id:      12,
		Address: "",
		Family:  "test",
	}

	/**
	高性能JSON特丹： jsoniter
	1、快，并且很快
	2、支持java和go
	3、100兼容JSON标准库，一行代码迁移到 jsoniter
	4、安装： go get github.com/json-iterator/go
	*/
	//json := jsoniter.ConfigCompatibleWithStandardLibrary//加上的话就是用json-iterator库

	data, err := json.Marshal(&u)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	//json_decode
	data1 := []byte(`{"name":"huangbin","id":"12","Family":"oneFamily"}`)
	u2 := &User{}

	err1 := json.Unmarshal(data1, u2)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(u2)      //&{huangbin 0 12  }，不会输出Family，因为struct已经声明不进行序列化和反序列化
	fmt.Println(u2.Name) //输出huangbin

}

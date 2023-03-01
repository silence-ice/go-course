package main

import "fmt"

/**
1、空接口可以保存任何类型这个特性可以方便地用于容器的设计，比如map类型可以将任意类型的值做成键值对保存，然后进行找回、遍历操作
*/

// 字典结构
type Dictionary struct {
	data map[interface{}]interface{} // 键值都为interface{}类型
}

// 结构方法——根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// 结构方法——设置键值
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// 结构方法——遍历所有的键值，如果回调返回值为false，停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// 结构方法——清空所有的数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// 普通方法——创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}
	// 初始化map
	d.Clear()
	return d
}

func main() {
	// 创建字典实例
	dict := NewDictionary()
	// 添加游戏数据
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	// 获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite) //favorite: 36

	// 遍历所有的字典元素，该传值是单个遍历方法，套到Visit就会循环整个字典
	dict.Visit(func(key, value interface{}) bool {
		// 将值转为int类型，并判断是否大于40
		if value.(int) > 40 {
			// 输出很贵
			fmt.Println(key, "is expensive")
			return true
		}
		// 默认都是输出很便宜
		fmt.Println(key, "is cheap")
		return true
	})
}

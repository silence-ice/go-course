package main

import "fmt"

func main() {
	//定义map，map就是无序hash集合
	//1、声明集合 var map_variable map[key_data_type]value_data_type
	//除了slice、map、function的内建类型都可以作为key
	person := map[string]string{
		"name": "huangbin",
		"age":  "23",
	}

	for one := range person {
		fmt.Println(person[one])
	}

	//2、使用 make 函数
	//map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity，格式如下：
	//map_variable := make(map[key_data_type]value_data_type, cap)
	countryCapitalMap := make(map[string]string) //创建集合

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		//Italy 首都是 罗马
		//Japan 首都是 东京
		//India  首都是 新德里
		//France 首都是 巴黎
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*使用键输出地图值 */
	for key, value := range countryCapitalMap {
		//Italy 首都是 罗马
		//Japan 首都是 东京
		//India  首都是 新德里
		//France 首都是 巴黎
		fmt.Println(key, value)
	}

	fmt.Println(countryCapitalMap["India"]) //新德里
	fmt.Println(countryCapitalMap["xxxx"])  //输出空，当key不存在的时候获取的是value的初始值，并不会报错

	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["India"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		//American 的首都是 新德里
		fmt.Println("India 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	//删除元素
	delete(countryCapitalMap, "France")

}

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()                    //获取当前时间
	fmt.Printf("current time:%v\n", now) //2020-02-13 11:40:05.444941 +0800 CST m=+0.000232501

	timestamp1 := now.Unix()                          //时间戳
	timestamp2 := now.UnixNano()                      //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1) //1581565252
	fmt.Printf("current timestamp2:%v\n", timestamp2) //1581565252277546000

	year := now.Year()                                  //年
	month := now.Month()                                //月
	day := now.Day()                                    //日
	hour := now.Hour()                                  //小时
	minute := now.Minute()                              //分钟
	second := now.Second()                              //秒
	fmt.Println(year, month, day, hour, minute, second) //2020 February 13 11 40 5

	//指定时间戳转时间
	timeObj := time.Unix(1581481569, 0)                       //将时间戳转为时间格式
	fmt.Println(timeObj)                                      //2020-02-13 11:41:42 +0800 CST
	year1 := timeObj.Year()                                   //年
	month1 := timeObj.Month()                                 //月
	day1 := timeObj.Day()                                     //日
	hour1 := timeObj.Hour()                                   //小时
	minute1 := timeObj.Minute()                               //分钟
	second1 := timeObj.Second()                               //秒
	fmt.Println(year1, month1, day1, hour1, minute1, second1) //2020 February 12 12 26 9

	//计算历史未来时间戳
	m, _ := time.ParseDuration("-1m")
	//m, _ := time.ParseDuration("-1h")
	fmt.Println(m) //-1m0s
	later := now.Add(m)
	fmt.Println(later.Unix()) //1分钟之前的时间戳

	//计算时间间隔
	before := now.Sub(later) // 当前时间和加1小时后的时间间隔
	fmt.Println(before)      //-1h0m0s

	/**
	Equal
	1、判断两个时间是否相同：func (t Time) Equal(u Time) bool
	2、Equal 函数会考虑时区的影响，因此不同时区标准的时间也可以正确比较，Equal 方法和用 t==u 不同，Equal 方法还会比较地点和时区信息。

	Before
	1、判断一个时间点是否在另一个时间点之前：func (t Time) Before(u Time) bool
	2、如果 t 代表的时间点在 u 之前，则返回真，否则返回假。

	After
	1、判断一个时间点是否在另一个时间点之后：func (t Time) After(u Time) bool
	2、如果 t 代表的时间点在 u 之后，则返回真，否则返回假。
	*/
	nowTime := time.Now() //获取当前时间
	minuteLaterExp, _ := time.ParseDuration("1m")
	minuteLaterTime := now.Add(minuteLaterExp) //一分钟后
	minuteBeforeExp, _ := time.ParseDuration("-1m")
	minuteBeforeTime := now.Add(minuteBeforeExp) //一分钟前

	fmt.Println(minuteBeforeTime.Before(nowTime)) //输出true
	fmt.Println(minuteLaterTime.After(nowTime))   //输出true

	//fmt.Println(nowTime)
	//fmt.Println(minuteLaterTime)
	//fmt.Println(minuteBeforeTime)

	/**
	 * time.Duration 时间段接受者的方法：输出和改变时间段单位。
	 *	func (d Duration) String() string // 格式化输出 Duration
	 *	func (d Duration) Nanoseconds() int64 // 将时间段表示为纳秒
	 *	func (d Duration) Seconds() float64 // 将时间段表示为秒
	 *	func (d Duration) Minutes() float64 // 将时间段表示为分钟
	 *	func (d Duration) Hours() float64 // 将时间段表示为小时
	 */
	fmt.Println(time.Duration(5000).String())           //5µs （微秒为单位）
	fmt.Println(time.Duration(5000))                    //5µs 功能同上，简写版
	fmt.Println(time.Millisecond)                       //1ms
	fmt.Println(time.Duration(5000) * time.Millisecond) //5µs * 1ms = 5s

	fmt.Println(time.Duration(5000).Nanoseconds())   //5000 （纳秒为单位）
	fmt.Println(time.Duration(5000000000).Seconds()) //5 （秒为单位）
}

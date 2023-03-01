package main

import "fmt"
import "time"

var (
	oneQuestion, twoQuestion, threeQuestion string
)

func main() {
	//one
	fmt.Printf("我们第一次交往的日子是什么时候:（格式：2021-05-20） ")
	fmt.Scanln(&oneQuestion) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。

	var oneAnswer bool = false

	for oneAnswer == false {
		if oneQuestion == "2017-12-25" {
			fmt.Println("恭喜你，答对啦！！")
			oneAnswer = true
		} else {
			fmt.Printf("错！再好好想想 ")
			fmt.Scanln(&oneQuestion) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
		}
	}

	//two
	fmt.Printf("我们婚礼的日子是什么时候:（格式：2021-05-20） ")
	fmt.Scanln(&twoQuestion) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。

	var twoAnswer bool = false

	for twoAnswer == false {
		if twoQuestion == "2020-12-24" {
			fmt.Println("恭喜你，答对啦！！")
			twoAnswer = true
		} else {
			fmt.Printf("错！再好好想想 ")
			fmt.Scanln(&twoQuestion) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
		}
	}

	//three
	fmt.Printf("我们的纪念信物是啥？ ")
	fmt.Scanln(&threeQuestion) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。

	var threeAnswer bool = false

	for threeAnswer == false {
		if threeQuestion == "香囊" {
			fmt.Println("很好！你通过了考验，打开支付宝，搜红包，输入密码 老婆520快乐呀 就可以领取红包啦")
			time.Sleep(time.Duration(60) * time.Second)
			threeAnswer = true
		} else {
			fmt.Printf("错！再好好想想 ")
			fmt.Scanln(&threeQuestion) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
		}
	}

	//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build love2.go

}

package main

import (
	"fmt"
)

func main() {
	//one 基本的实例化形式
	/**
	结构体的定义格式如下：
		type 类型名 struct {
			字段1 字段1类型
			字段2 字段2类型
			…
		}
	*/
	type Point struct {
		X int
		Y int
	}

	//实例化Books实例，以 var 的方式声明结构体即可完成实例化
	var p Point //Point 为结构体类型，p 为结构体的实例
	p.X = 10    //使用.来访问结构体的成员变量，如p.X和p.Y等，结构体成员变量的赋值方法与普通变量一致
	p.Y = 20
	fmt.Println(p)   //{10 20}
	fmt.Println(p.X) //{10 20}

	//two 创建指针类型的结构体
	type Player struct {
		Name        string
		HealthPoint int
		MagicPoint  int
	}

	player := new(Player) //Player 为结构体类型，Player 类型被实例化后保存到 ins 变量中，player 的类型为 *Player，属于指针。
	player.Name = "HuangBin"
	player.HealthPoint = 100
	player.MagicPoint = 200
	fmt.Println(player)             //&{HuangBin 100 200}
	fmt.Println(player.HealthPoint) //100

	//对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
	player2 := &Player{}
	player2.Name = "LiuHuiLing"
	player2.HealthPoint = 200
	player2.MagicPoint = 400
	fmt.Println(player2)             //&{LiuHuiLing 200 400}
	fmt.Println(player2.HealthPoint) //200

	player3 := &Player{
		Name:        "huangbin",
		HealthPoint: 100,
		MagicPoint:  100,
	}
	fmt.Println(player3)             //&{LiuHuiLing 200 400}
	fmt.Println(player3.HealthPoint) //200
}

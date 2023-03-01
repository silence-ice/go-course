package _2_goTest

import (
	"fmt"
	"testing"
)

func print3(t *testing.T) {
	result := calNum2(5)
	t.Log(result)
}

func print4(t *testing.T) {
	result := calNum2(10)
	t.Log(result)
}

func TestPrintAll2(t *testing.T) {
	t.Run("testPrint3", print3)
	t.Run("testPrint4", print4)
}

// test脚本入口
func TestMain(m *testing.M) {
	//测试案例的时候可能需要初始化数据库连接，文件打开，REST服务器登录等，可以使用TestMain作为初始化test,并且使用m.Run()来调用其他tests

	fmt.Println("初始化数据库链接")
	m.Run() //如果没有在TestMain中调用m.Run()则除了TestMain以外的其他tests都不会被执行
}

func calNum2(num int) int {
	sum := 0
	i := 1
	for ; i < num; i++ {
		sum = sum + i
	}

	return sum
}

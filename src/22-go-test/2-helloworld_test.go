package _2_goTest

import (
	"testing"
)

func calNum(num int) int {
	sum := 0
	i := 1
	for ; i < num; i++ {
		sum = sum + i
	}

	return sum
}

func print1(t *testing.T) {
	result := calNum(5)
	t.Log(result)
}

func print2(t *testing.T) {
	result := calNum(10)
	t.Log(result)
}

func TestPrintAll(t *testing.T) {
	//go 的test不会保证多个Testxxx方法是按顺序执行，所以可以使用 t.Run 来执行 子test函数 来控制test输出以及test的顺序

	t.Run("testPrint1", print1)
	t.Run("testPrint2", print2)
}

/**
命令行执行： silence$go test -v 2-helloworld_test.go
			=== RUN   TestAll
			=== RUN   TestAll/testPrint1
			=== RUN   TestAll/testPrint2

			--- PASS: TestAll (0.00s)
				--- PASS: TestAll/testPrint1 (0.00s)
					2-helloworld_test.go:19: 10
				--- PASS: TestAll/testPrint2 (0.00s)
					2-helloworld_test.go:25: 45
			PASS
			ok      command-line-arguments  0.005s
*/

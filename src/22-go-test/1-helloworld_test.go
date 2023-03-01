package _2_goTest

import (
	"fmt"
	"testing"
)

/*
*
什么是单元测试？
1、单元测试是在软件开发过程中要进行的最低级别的测试活动，软件的独立单元将在与程序的其他部分相隔离的情况下测试。
2、单元测试是指对软件中的最小可测试单元进行检查和验证，单元就是人为规定的最小的被测功能模块，可以是函数也可以是类或者是一个小功能模块。

单元测试规则：
1、【文件命名规则】在命名文件时需要让文件必须以_test结尾
2、【函数命名规则】单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以Test为前缀，例如：

	func TestXXX( t *testing.T )
	2.1 测试用例文件不会参与正常源码编译，不会被包含到可执行文件中。
	2.2 测试用例文件不需要 main() 作为函数入口。所有在以_test结尾的源码内以Test开头的函数会自动被执行。
	2.3 测试用例可以不传入 *testing.T 参数。

3、【命令执行规则】默认的情况下，go test 命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数：

	-run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数
	-v 显示测试的详细命令
	-cover 开启测试覆盖率

	go test -v helloworld_test.go => 表示运行helloworld_test.go里面所有的测试用例
	go test -v -run TestA helloworld_test.go => 表示运行helloworld_test.go里面所有以TestA开头的测试用例，比如TestAB测试用例也会被执行
	go test -v -run TestA$ helloworld_test.go => 表示只执行 TestA 测试用例。

单元测试日志：
每个测试用例可能并发执行，使用 testing.T 提供的日志输出可以保证日志跟随这个测试上下文一起打印输出。testing.T 提供了几种日志输出方法

	Log		打印日志，同时结束测试
	Logf	格式化打印日志，同时结束测试
	Error	打印错误日志，同时结束测试
	Errorf	格式化打印错误日志，同时结束测试
	Fatal	打印致命日志，同时结束测试
	Fatalf	格式化打印致命日志，同时结束测试
*/
func TestHelloworld(t *testing.T) {
	fmt.Println("TestHelloworld")
	t.Log("hello world")
}

func TestHelloworld2(t *testing.T) {
	//t.kipNow() 作用是跳过当前test，并且直接PASS处理下一个test，这必须放在某个test case的第一行才起作用
	t.SkipNow()

	fmt.Println("TestHelloworld2")
	t.Log("hello world2")
}

/**
命令行执行： silence$go test -v 1-helloworld_test.go    	=> 在 go test 后跟 1-helloworld_test.go 文件，表示测试这个文件里的所有测试用例。

			=== RUN   TestHelloworld				=> 表示开始运行名叫 TestHelloWorld 的测试用例。
			huangbin								=> 打印字符串 huangbin。
			--- PASS: TestHelloworld (0.00s)		=> 表示已经运行完 TestHelloWorld 的测试用例，PASS 表示测试成功。
				helloworld_test.go:39: hello world	=> 打印字符串 hello world。


			=== RUN   TestHelloworld2				=> 表示开始运行名叫 TestHelloWorld2 的测试用例。
			--- SKIP: TestHelloworld2 (0.00s)		=> 表示跳过 TestHelloworld2 的测试用例

			PASS
			ok      command-line-arguments  0.005s	=> 显示测试结果，ok 表示测试通过，command-line-arguments 是测试用例需要用到的一个包名，0.005s 表示测试花费的时间。
*/

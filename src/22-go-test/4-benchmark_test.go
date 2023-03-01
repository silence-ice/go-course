package _2_goTest

import (
	"testing"
)

/*
*
什么是基准测试？
1、基准测试可以测试一段程序的运行性能及耗费 CPU 的程度。Go语言中提供了基准测试框架，使用方法类似于单元测试。

基准测试规则：
1、【文件命名规则】基准测试的代码文件必须以_test.go结尾
2、【函数命名规则】单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以Benchmark为前缀，例如：

	func Benchmark_XXX(t *testing.B)
	2.1 基准测试的函数必须以Benchmark开头，必须是可导出的
	2.2 基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
	2.3 基准测试函数不能有返回值
	2.4 b.N是基准测试框架提供的，表示for循环的次数，在执行过程中会根据每个样本的执行时间是否稳定， 来决策是否增加b.N的次数 以达到稳态
	2.5 b.ResetTimer是重置计时器，这样可以避免for 循环之前的初始化代码的干扰

3、【命令执行规则】默认的情况下，go test 命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数：

	-v 				显示测试的详细命令
	-bench regexp 	执行相应的 benchmarks
		-bench=. 	 表示运行 benchmark_test.go 文件里的所有基准测试，和单元测试中的-run类似。
		-bench=Alloc 表示指定只测试 Benchmark_Alloc() 函数。
	-benchmem		以显示内存分配情况
	-benchtime=5s	通过-benchtime参数可以自定义测试时间为5s

	go test -v -bench=. benchmark_test.go => 表示运行 benchmark_test.go里面所有的基准测试用例
	go test -v -bench=Add benchmark_test.go => benchmark_test.go里面所有以Benchmark_Add的基准测试用例

4、基准测试原理

	4.1 基准测试框架对一个测试用例的默认测试时间是 1 秒。
	4.2 开始测试时，当以 Benchmark 开头的基准测试用例函数返回时还不到 1 秒，那么 testing.B 中的 N 值将按 1、2、5、10、20、50……递增，同时以递增后的值重新调用基准测试用例函数。
*/
func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
	/**
	  silence$go test -v -bench=Add benchmark_test.go
	  goos: darwin										=> 系统参数
	  goarch: amd64										=> 处理器类型参数
	  Benchmark_Add-4  2000000000  0.58 ns/op			=> Benchmark_Add-4表示Benchmark名字-CPU；进行了500次循环测试，0.58平均每次执行时间
	  PASS
	  ok      command-line-arguments  1.233s			=> 显示测试结果，ok 表示测试通过，command-line-arguments 是测试用例需要用到的一个包名，1.233s 表示测试花费的时间。

	  silence$go test -v -bench=Add -benchmem benchmark_test.go
		goos: darwin
		goarch: amd64
		Benchmark_Add-4 2000000000 0.57 ns/op 16 B/op 2 allocs/op => “16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。
		PASS
		ok      command-line-arguments  1.211s
	*/
}

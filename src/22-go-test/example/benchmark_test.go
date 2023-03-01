package example

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"testing"
)

/*
*
测试规则：
1、如果测试单个文件，一定要带上被测试的原文件，如果原文件有其他引用，也需一并带上，如果该测试文件如果有引用其他文件就会报错：

	比如： undefined: GoodsSigned

2、执行 $go test -v -bench=. benchmark_test.go 报错原因:

	2.1 go test与其他的指定源码文件进行编译或运行的命令程序一样会为指定的源码文件生成一个虚拟代码包——“command-line-arguments”
	2.2 对于运行这次测试的命令程序来说，测试源码文件 benchmark_test.go是属于代码包“command-line-arguments”的
	2.3 可是它引用了其他包中的数据并不属于代码包“command-line-arguments”，编译不通过，错误自然发生了。

3、解决方法:

	3.1 在当前整个目录运行
		go test -v -bench=.
	3.2 执行命令时加入这个测试文件需要引用的源码文件，在命令行后方追加的文件都会被加载到command-line-arguments中进行编译
		go test -v -bench=UpdateForLock benchmark_test.go mysql.go
*/
func Benchmark_UpdateForLock(b *testing.B) {
	//先插入一条资源，因为over.go已经声明了db变量为mysql资源链接了，整个项目共享变量，所以无需声明
	g := GoodsSigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(1000000)
	_, err := db.Insert(g)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		//fmt.Println("count:",i)
		UpdateForLock(g.Goods)
	}
}

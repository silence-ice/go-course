package _2_goTest

import "testing"

// 该方法会 使每个样本的执行时间 随着b.N增加而增加，从而失去稳定，促使增加b.N的次数 以达到稳态
func makeAdd(n int) int {
	for n > 0 {
		n--
	}
	return n
}

// go test -v -bench=TestN 5-benchmark_test.go
func Benchmark_TestN(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		makeAdd(n)
	}

}

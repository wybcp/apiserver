package util

import (
	"testing"
)

func TestGenShortId(t *testing.T) {
	shortID, err := GenShortId()
	if shortID == "" || err != nil {
		t.Error("GenShortId failed")
	}
	t.Log("GenShortId test pass")
}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	// 在 b.StopTimer() 和 b.StartTimer() 之间可以做一些准备工作，这样这些时间不影响我们测试函数本身的性能。
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortID, err := GenShortId()
	if shortID == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

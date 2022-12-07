package allocation

import "testing"

func BenchmarkNewByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewDataByValue()
	}
}

func BenchmarkNewByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewDataByPointer()
	}
}

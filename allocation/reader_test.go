package allocation

import "testing"

func BenchmarkSliceWithArgument(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := make([]byte, bufferSize)
		reader := ReaderWithSliceArgument{}
		_, _ = reader.Read(p)
	}
}

func BenchmarkSliceWithReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := ReaderWithSliceReturn{}
		_, _ = reader.Read(bufferSize)
	}
}

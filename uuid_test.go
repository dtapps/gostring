package gostring

import (
	"testing"
)

func TestGetUuId(t *testing.T) {
	t.Log(GetUuId())
}

func BenchmarkGetUuId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetUuId())
	}
}

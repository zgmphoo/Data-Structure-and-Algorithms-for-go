package main

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var slice = []int{44, 80, 78, 77, 42}
		BubbleSort(slice)
	}
}

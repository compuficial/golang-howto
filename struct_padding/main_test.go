package main

import "testing"

const N = 1000000

var (
	inefficientSlice = make([]Inefficient, N)
	efficientSlice   = make([]Efficient, N)
)

var sink int64

func BenchmarkIterateInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sum int64
		for j := range inefficientSlice {
			sum += inefficientSlice[j].field2 + inefficientSlice[j].field4
		}
		sink = sum
	}
}

func BenchmarkIterateEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sum int64
		for j := range efficientSlice {
			sum += efficientSlice[j].field2 + efficientSlice[j].field4
		}
		sink = sum
	}
}

func BenchmarkAllocateInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]Inefficient, 10000)
	}
}

func BenchmarkAllocateEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]Efficient, 10000)
	}
}

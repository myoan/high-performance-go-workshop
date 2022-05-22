package main

import (
	"testing"
	"time"
)

func BenchmarkFib3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(3)
	}
}

func boringAndExpensiveSetup() {
	time.Sleep(100 * time.Millisecond)
}

func BenchmarkFib20(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Fib(20)
	}
}

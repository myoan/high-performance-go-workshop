package main

import "testing"

var v = make([]int, 9)

// var v [9]int
var A, B, C, D, E, F, G, H, I int

func BenchmarkBoundsCheckInOrder(b *testing.B) {
	// var v = make([]int, 9)
	for n := 0; n < b.N; n++ {
		// A = v[0]
		// B = v[1]
		// C = v[2]
		// D = v[3]
		// E = v[4]
		// F = v[5]
		// G = v[6]
		// H = v[7]
		// I = v[8]
		A = v[8]
		B = v[0]
		C = v[1]
		D = v[2]
		E = v[3]
		F = v[4]
		G = v[5]
		H = v[6]
		I = v[7]
	}
}

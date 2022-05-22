package main

func Fib(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	default:
		return Fib(n-2) + Fib(n-1)
	}
}

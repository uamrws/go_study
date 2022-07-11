package fibo

import "testing"

func benchmarkFibonacci(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}
func benchmarkFibonacci2(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fibonacci2(n)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	benchmarkFibonacci(b, 10)
	b.Log(Fibonacci(10))
}
func BenchmarkFibonacci_1(b *testing.B) {
	benchmarkFibonacci(b, 20)
	b.Log(Fibonacci(20))
}
func BenchmarkFibonacci_2(b *testing.B) {
	benchmarkFibonacci(b, 50)
	b.Log(Fibonacci(50))
}

func BenchmarkFibonacci2(b *testing.B) {
	benchmarkFibonacci2(b, 10)
	b.Log(Fibonacci2(10))
}
func BenchmarkFibonacci2_1(b *testing.B) {
	benchmarkFibonacci2(b, 20)
	b.Log(Fibonacci2(20))
}
func BenchmarkFibonacci2_2(b *testing.B) {
	benchmarkFibonacci2(b, 50)
	b.Log(Fibonacci2(50))
}

package fibo

func fibonacci() fibGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fibGen func() int

func Fibonacci(n int) int {
	fibgen := fibonacci()
	result := 0
	for i := 0; i < n; i++ {
		result = fibgen()
	}
	return result
}

func fibonacci2(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci2(n-2) + fibonacci2(n-1)
}

func Fibonacci2(n int) int {
	return fibonacci2(n)
}

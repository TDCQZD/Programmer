package src

// 递归, 间复杂度O(2^n) 不推荐
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//循环实现 时间复杂度O(n) 推荐
func Fibonacci1(n int) int {
	array := []int{0, 1}
	if n < 2 {
		return array[n]
	}

	fibMinusOne := 1
	fibMinusTwo := 0
	fibN := 0
	i := 2
	for i <= n {
		fibN = fibMinusOne + fibMinusTwo
		fibMinusTwo = fibMinusOne
		fibMinusOne = fibN
		i++
	}
	return fibN
}


package src

//Fibonacci_book 循环实现 时间复杂度O(n) 推荐
func Fibonacci_book(n int) int {
	array := make([]int, n+1)

	array[0] = 1
	array[1] = 1
	i := 2
	for {
		array[i] = array[i-1] + array[i-2]
		i++
		if i > n {
			break
		}
	}
	// for i <= n {
	// 	array[i] = array[i-1] + array[i-2]
	// 	i++
	// }

	return array[n]
}

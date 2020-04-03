package src

//可能引起死循环的解法
func NumberOf1(n int) int {
	count := 0

	for {
		if n&1 == 1 {
			count++
		}
		n = n >> 1

		if n == 0 {
			break
		}
	}
	return count
}

// 常规解法:go中无法实现
func NumberOf1_1(n int) int {
	count := 0
	flag := 1
	for {
		/*
			注意:参考C代码，在C中非0无符号整数是true,即 n&flag[int] true
			go中 type int not used as if condition
		*/
		if n&flag > 0 {
			count++
		}
		flag = flag << 1

		if flag == 0 {
			break
		}
	}
	return count
}

// 推荐解法
func NumberOf1_2(n int) int {
	count := 0
	for {
		count++
		n = (n - 1) & n
		if n == 0 {
			break
		}
	}
	return count
}

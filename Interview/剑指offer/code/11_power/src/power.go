package src

func PowerWithExponent(base float64, exponent int) (res float64) {

	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	// 位移运算代替除以2
	res = PowerWithExponent(base, exponent>>1)

	res *= res
	// 位于运算符&代替求余运算符%，来判断一个数是奇数或者偶数
	if exponent&0x1 == 1 {
		res *= base
	}
	return res
}

package src

// 字符串装数字的工具方法
func StringToint(str string) int {
	b := []byte(str) //把字符串转换为字节数组
	result := 0
	for i := 0; i < len(b); i++ {
		if b[i] >= 48 && b[i] <= 57 {
			result = result*10 + (int(b[i]) - 48)
		}
	}
	return result
}

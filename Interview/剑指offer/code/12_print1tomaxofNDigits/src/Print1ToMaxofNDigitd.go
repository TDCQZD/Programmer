package src

import (
	"strconv"
	"fmt"
)

// 未考虑大数问题
func Print1ToMaxofNDigitd_1(num int) {
	number := 1

	for i := 1; i < num; i++ {
		number *= 10
	}
	for i := 1; i < number; i++ {
		fmt.Println(i)
	}
}

/*

 */
func Print1ToMaxofNDigitd_2(num int) {
	if num <= 0 {
		return
	}
	number := make([]string, num+1)
	// 字符串中的毎一个数字都初始化为'0'
	for i := 0; i < len(number); i++ {
		number[i] = "0"
	}

	// fmt.Println(number)
	for !Increment(number) { //把每一个字符串表示的数字+1
		PrintNumber(number) //打印字符串数字
	}
}

// 字符串表达的数字上模拟加法:把每一个字符串表示的数字+1
func Increment(number []string) bool {
	nLength := len(number) - 1
	for number[nLength] == "9" { // 从字符的最右边开始，如果为‘9’就置‘0’，直到不是‘9’就将该位加1
		number[nLength] = "0"
		nLength--
	}
	// number[nLength] += 1
	if number[0] == "0" {
		return true
	}
	return false

}
// 大数相加
func AddStrings(num1 string, num2 string) string {
    len1 := len(num1)
    len2 := len(num2)
    var ret string
    var carry int
    var digit int

    for i, j := len1, len2; i > 0 || j > 0 || carry > 0; {
        digit = carry
        if i > 0 {
            i--
            digit += int(num1[i] - '0')
        }
        if j > 0 {
            j--
            digit += int(num2[j] - '0')
        }

        if digit > 9 {
            carry = 1
        } else {
            carry = 0
        }

        ret = strconv.Itoa(digit % 10) + ret
    }

    return ret
}

func PrintNumber(number []string) {
	isBeginning0 := true
	nLength := len(number)

	for i := 0; i < nLength; i++ {
		if isBeginning0 && number[i] != "0" {
			isBeginning0 = false
		}
		if !isBeginning0 {
			fmt.Println(number[i])
		}
	}
}

func Print1ToMaxofNDigitd_3(num int) {
	if num <= 0 {
		return
	}
	number := []int{}
	for i := 0; i < 10; i++ {
		number[i] = i
	}
}

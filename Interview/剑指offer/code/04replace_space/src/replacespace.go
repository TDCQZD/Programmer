package src

import (
	"fmt"
	"strings"
)

/*ReplaceSpace 替换字符串中空格*/
func ReplaceSpaceByInner(str string) string {
	r := strings.NewReplacer(" ", "%20")
	return r.Replace("We are happy.")
}

func ReplaceSpace(str string) string {
	if str == " " && len(str) <= 0 {
		return ""
	}
	originalLength := len(str) //字符串str实际长度
	numberOfBlank := 0         //字符串str中 空格数目
	// 遍历获取字符串str中空格个数
	s := " "
	for _, value := range str {
		s = fmt.Sprintf("%c", value)
		if s == " " {
			numberOfBlank++
		}
	}
	
	newLength := originalLength + numberOfBlank
	
	newStr := make([]string, newLength)
	
	count := 0
	for _, value := range str {
		s = fmt.Sprintf("%c", value)
		if s == " " {
			newStr = append(newStr, "%")
			newStr = append(newStr, "2")
			newStr = append(newStr, "0")
		} else {
			newStr = append(newStr, s)
		}
		count++
	}
	fmt.Println(count)
	fmt.Println(newStr)
	fmt.Println(len(newStr))
	// str = ""
	// for _, value := range newStr {
	// 	s = fmt.Sprintf("%c", value)
	// 	str = str + s
	// }
	return str
}

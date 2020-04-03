package src

import (
	"fmt"
	"testing"
)

func TestNumberOf1(t *testing.T) {
	res := 0
	res=NumberOf1_1(3)
	fmt.Println(res)
	res=NumberOf1_1(8)
	fmt.Println(res)
	res=NumberOf1_1(11)
	fmt.Println(res)
	
}

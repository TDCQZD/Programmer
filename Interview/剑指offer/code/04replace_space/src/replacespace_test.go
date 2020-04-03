package src

import (
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	// testInner(t);
	testArray(t);

}
func testArray(t *testing.T) {
	//ReplaceSpace
	str := "We are happy."
	
	str = ReplaceSpace(str)
	fmt.Println(str)
	
}
func testInner(t *testing.T) {
	// go 内置方法实现
	str := "We are happy."
	str = ReplaceSpaceByInner(str)
	fmt.Println(str)
}

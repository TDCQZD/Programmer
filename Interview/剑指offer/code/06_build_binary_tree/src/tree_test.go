package src

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 8, 6}
	post := []int{7, 4, 2, 5, 8, 6, 3, 1}
	// tree := buildTreeByPre(pre, in)
	tree := buildTreeByPost(in, post)
	fmt.Println(pre)
	// PreOrder(tree)
	fmt.Println("--------------------------")
	// InfixOrder(tree)
	fmt.Println("--------------------------")
	PostOrder(tree)
}

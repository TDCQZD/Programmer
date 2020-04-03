package src

import (
	"testing"
)

func TestCopyComplexList(t *testing.T) {
	pNode1 := CreateNode(1)
	pNode2 := CreateNode(2)
	pNode3 := CreateNode(3)
	pNode4 := CreateNode(4)
	pNode5 := CreateNode(5)

	BuildNodes(pNode1, pNode2, pNode3)
	BuildNodes(pNode2, pNode3, pNode5)
	BuildNodes(pNode3, pNode4, nil)
	BuildNodes(pNode4, pNode5, pNode2)

	res := CopyComplexList(pNode1)
	PrintList(res)
}

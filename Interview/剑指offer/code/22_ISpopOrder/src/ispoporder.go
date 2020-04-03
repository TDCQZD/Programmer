package src

import "fmt"

func IsPopOrder(pPush, pPop []int, nLength int) (flag bool) {
	flag = false
	if pPush != nil && pPop != nil && nLength > 0 {
		pNextPush := 0
		pNextPop := 0
		stack := NewStack()

		for {
			for {
				if pNextPush == nLength-1 || stack.Top() == pPop[pNextPop] {
					break
				}
				fmt.Println("push:", pPush[pNextPush])
				stack.Push(pPush[pNextPush])
				pNextPush++
			}

			if pNextPop >= nLength || stack.Top() != pPop[pNextPop] {
				break
			}
			res := stack.Pop()
			pNextPop++
			fmt.Println("pop:", res)
		}

		if stack.IsEmpty() && pNextPop == nLength-1 {
			flag = true
		}

	}
	return
}

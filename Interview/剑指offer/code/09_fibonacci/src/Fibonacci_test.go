package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	// testFibonacci(t, 45)
	// testFibonacci1(t, 1)
	testFibonacci_book(t,45)
}
func testFibonacci(t *testing.T, n int) {
	res := Fibonacci(n)
	fmt.Println("O(2^n):", res)
}
func testFibonacci1(t *testing.T, n int) {
	res := Fibonacci1(n)
	fmt.Println("O(n):", res)
}

func testFibonacci_book(t *testing.T, n int) {
	res := Fibonacci1(n)
	fmt.Println("book", res)
}

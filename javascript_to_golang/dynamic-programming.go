package main

import (
	"fmt"
)

var counterPlainRecursive = 0
var counterDynamicProgramming = 0

func fibonacciPlainRecursive(n int) int {
	counterPlainRecursive++
	if n < 2 {
		return n
	}
	return fibonacciPlainRecursive(n-1) + fibonacciPlainRecursive(n-2)
}

func fibonacciDynamicProgramming() func(int) int {
	cache := make(map[int]int)
	var fib func(int) int
	fib = func(n int) int {
		counterDynamicProgramming++
		if val, ok := cache[n]; ok {
			return val
		}
		if n < 2 {
			return n
		}
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}
	return fib
}

func main() {
	fibonacciPlainRecursive(20)
	fasterFibonacci := fibonacciDynamicProgramming()
	fasterFibonacci(20)
	fmt.Println("we did", counterPlainRecursive, "calculations for Plain Recursive")
	fmt.Println("we did", counterDynamicProgramming, "calculations for Dynamic Programming")
}

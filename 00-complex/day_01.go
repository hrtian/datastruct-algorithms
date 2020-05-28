package main

import (
	"fmt"
	"time"
)

func check(title string, n int, f func(int) int) {
	if f == nil {
		return
	}

	if title == "" {
		title = ""
	} else {
		title = "【" + title + "】"
	}
	fmt.Printf("%s\nStart:\n", title)
	fmt.Printf("Start time: %v\n", time.Now().Format("15:04:05.000"))
	start := time.Now()
	f(n)
	end := time.Now()
	fmt.Printf("End time: %v\n", end.Format("15:04:05.000"))
	fmt.Printf("Duration: %s\n", end.Sub(start))
}

func main() {
	n := 32
	check("Fibonacci 01", n, fibonacci1)
	check("Fibonacci 02", n, fibonacci)
}

// this is a program for fibonacci numbers
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci1(n int) int {
	if n <= 1 {
		return n
	}
	// 0 1 1 2 3 5
	first := 0
	second := 1
	for i := 0; i < n; i++ {
		sum := first + second
		first = second
		second = sum
	}
	return first
}

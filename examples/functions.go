package examples

import "fmt"

func recursionExample() {

	// forward declaration needed because of recursion
	var fib func(int) int

	fib = func(n int) int {
		if n < 2 {
			return 1
		}
		return fib(n-2) + fib(n-1)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}
}

func multReturnExample() {

	div := func(n int, d int) (int, int) {
		return n / d, n % d
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 8; j++ {
			q, r := div(i*i, j)
			fmt.Printf("%-2d/%-2d=%-2d,%-2d ", i*i, j, q, r)
		}
		fmt.Println()
	}
}

func variadicExample() {

	max := func(list ...int) (result int) {
		result = list[0]
		for _, n := range list {
			if result < n {
				result = n
			}
		}
		return
	}

	fmt.Println("max =>", max(3, 1, 4, 1, 5, 9))
	someNumbers := []int{3, 2, 3, 1, 2, 8, 5, 1}
	fmt.Println("max =>", max(someNumbers...))
}

func logger(prefix string) func(string) {
	return func(msg string) {
		fmt.Println(prefix + ": " + msg)
	}
}

func closureExample() {
	info := logger("information")
	warn := logger("warning")
	info("closure test")
	warn("dange zone")
}

// FuncExamples contains examples of functions in go
func FuncExamples() {
	fmt.Println("\nFunctions in go")
	fmt.Println("===============")

	recursionExample()
	multReturnExample()
	variadicExample()
	closureExample()
}

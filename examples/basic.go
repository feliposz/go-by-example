package examples

import (
	"fmt"
	"math"
	"time"
)

func helloExample() {
	fmt.Println("Hello, go!")
}

func valuesExample() {
	fmt.Println("Pi is", math.Pi)
}

func variablesExample() {
	var a = "simple"
	var b int
	b = 123
	c := 45.67
	d := float64(b) + c
	fmt.Println(a, b, c, d)
}

func constantsExample() {
	const myPi = 3.14149
	fmt.Println("In-house pi is", myPi)
}

func forExample() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d*%d=%-2d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func ifExample() {
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func switchExample() {

	var d string
	switch time.Now().Weekday() {
	case time.Sunday:
		d = "domenica"
	case time.Monday:
		d = "lunedì"
	case time.Tuesday:
		d = "martedì"
	case time.Wednesday:
		d = "mercoledì"
	case time.Thursday:
		d = "giovedì"
	case time.Friday:
		d = "venedì"
	case time.Saturday:
		d = "sabato"
	default:
		d = "?!?"
	}
	fmt.Printf("Ciao! Oggi è %s!\n", d)

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buon giorno")
	case t.Hour() < 18:
		fmt.Println("Buon pomeriggio")
	case t.Hour() < 22:
		fmt.Println("Buona sera")
	default:
		fmt.Println("Buona notte")
	}

}

func arraysExample() {
	var fib [10]int
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < len(fib); i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}
	for index, value := range fib {
		fmt.Printf("fib[%d] = %d\n", index, value)
	}
}

func slicesExample() {
	s := make([]string, 3)
	s[0] = "Apple"
	s[1] = "Google"
	s[2] = "Microsoft"

	s = append(s, "Amazon", "Facebook")

	// Remove first and last
	s = s[1 : len(s)-1]

	for index, value := range s {
		fmt.Printf("%d = %s\n", index, value)
	}
}

func mapExample() {
	port2ita := map[string]string{
		"eu":   "io",
		"tu":   "tu",
		"ele":  "lui",
		"nós":  "noi",
		"vós":  "voi",
		"eles": "loro",
	}

	for port, ita := range port2ita {
		fmt.Printf("\"%s\" è \"%s\" in italiano.\n", port, ita)
	}
}

func rangeExample() {
	odds := [...]int{1, 3, 5, 7, 9}
	evens := [...]int{2, 4, 6, 8, 10}
	all := make([]int, len(odds)+len(evens))

	for i, n := range odds {
		all[i*2] = n
	}
	for i, n := range evens {
		all[i*2+1] = n
	}
	fmt.Println(all)
}

// BasicExamples contains examples of basic functionality in go
func BasicExamples() {
	fmt.Println("\nSome basic functionality in go")
	fmt.Println("==============================")

	helloExample()
	valuesExample()
	variablesExample()
	constantsExample()
	forExample()
	ifExample()
	switchExample()
	arraysExample()
	slicesExample()
	mapExample()
	rangeExample()
}

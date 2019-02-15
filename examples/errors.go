package examples

import (
	"errors"
	"fmt"
)

func simpleErrorExample() {

	failTest := func() (string, error) {
		return "some value", errors.New("Simple error")
	}

	result, err := failTest()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

type customError struct {
	someErrorCode    int
	someErrorMessage string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.someErrorCode, e.someErrorMessage)
}

func customErrorExample() {

	failTest := func() (string, error) {
		return "some stuff", &customError{42, "don't panic"}
	}

	result, err := failTest()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func ErrorExamples() {
	fmt.Println("\nError handling")
	fmt.Println("==============")

	simpleErrorExample()
	customErrorExample()
}

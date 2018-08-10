package main

import "errors"
import "fmt"

// By convention, errors are the last return value and have type error,
// a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// Custom error types by using the Error() method on them
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// In this case we use &argError syntax to build a new struct, supply values
// for the two fields arg and prob.
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// There is an array init here that we loop through
	// this loops over 7 and then 42.
	for _, i := range []int{7, 42} {
		// Note the inline error check on if e != nil
		// This is common idiom in Go code.
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to programmatically use the data in a custom error, you'll
	// need to get the error as an instance of the custom error type via type
	// assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

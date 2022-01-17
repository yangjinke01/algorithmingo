package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg     int
	message string
}

func (a argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, message: "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, item := range []int{7, 42} {
		if result, err := f1(item); err != nil {
			fmt.Println("f1 failed:", err)
		} else {
			fmt.Println("f1 worked:", result)
		}
	}
	for _, item := range []int{7, 42} {
		if result, err := f2(item); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", result)
		}
	}
	//get the error as an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	}
}

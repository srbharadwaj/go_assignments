package main

import (
	"errors"
	"fmt"
)

var DivideByZero = errors.New("DivideByZero error occured")

func main() {

	divide(3, 1)
}

func checkDivideByZero(y int) error {
	if y == 0 {
		return DivideByZero
	}
	return nil
}
func divide(x, y int) {
	val := checkDivideByZero(y)
	if val == nil {
		quotient := x / y
		remainder := x % y
		fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", x, y, quotient, remainder)
	} else {
		fmt.Printf("error found!!! %v\n", val)
	}
}

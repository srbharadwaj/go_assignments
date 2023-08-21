package main

import "fmt"

var dummy = 100

func main() {
	/*fmt.Println(dummy)
	print("hello world!!\n")
	fmt.Println("hello from println")

	var name string
	name = "asdasd"
	fmt.Printf("this is a %s string\n ", name)

	var flag = true
	fmt.Println(flag)

	anotherstr := "some string" // only in func scope
	fmt.Println(anotherstr)

	var (
		x, y, str = 200, 100, "sum of %d and %d is %d\n"
		ressult   = x + y
	)
	fmt.Printf(str, x, y, ressult)

	x1, y1, str1 := 201, 101, "sum of %d and %d is %d\n"
	result1 := x1 + y1

	fmt.Printf(str1, x1, y1, result1)
	*/

I_LOOP:
	for i := 2; i <= 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				//fmt.Printf("%d is not a prime number\n", i)
				continue I_LOOP
			}
		}
		fmt.Printf("%d is a prime number\n", i)
	}
}

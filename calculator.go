package main

import "fmt"

func main() {
	var option int
	var num1, num2 int
	displaystr := `Enter any of the below options
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit
	`
	asknumber := "Please enter 2 values"

START:
	fmt.Println(displaystr)
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Println(asknumber)
		fmt.Scanln(&num1, &num2)
		fmt.Printf("Value is %d\n", num1+num2)
		goto START
	case 2:
		fmt.Println(asknumber)
		fmt.Scanln(&num1, &num2)
		fmt.Printf("Value is %d\n", num1-num2)
		goto START
	case 3:
		fmt.Println(asknumber)
		fmt.Scanln(&num1, &num2)
		fmt.Printf("Value is %d\n", num1*num2)
		goto START
	case 4:
		fmt.Println(asknumber)
		fmt.Scanln(&num1, &num2)
		fmt.Printf("Value is %d\n", num1/num2)
		goto START
	case 5:
		goto END
	default:
		fmt.Println("Invalid options")
		goto START
	}

END:
	fmt.Println("Thanks. Done")
}

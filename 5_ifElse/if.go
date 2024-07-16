package main

import "fmt"

func main() {
	// ** if else
	/*
		age := 17
		if age == 18 {
			fmt.Println("Person is adult")
		} else {
			fmt.Println("Person is not adult")
		}
	*/
	// ** else if
	/*
		age := 11

		if age >= 18 {
			fmt.Println("You are adult")
		} else if age >= 12 {
			fmt.Println("You Are teenager")

		} else {
			fmt.Println("You Are Kid")

		}
	*/
	// ** Logical operators in if-else
	const role string = "Admin"
	const hasPermissions bool = true
	if role == "Admin" && hasPermissions {
		// fmt.Println("Yes")
	}

	//** variable declaration inside if-else
	if age := 18; age >= 18 {
		fmt.Println("Your Age is ", age)
	} else {
		fmt.Println("Your Age is ", age)

	}
	// TODO: Go doesn't have ternary operator
}

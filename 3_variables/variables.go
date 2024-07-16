package main

import "fmt"

func main() {
	var hello string = "hello world"
	hello = "Hello there how are you doing?"
	fmt.Println(hello)
	// another way to declare variable
	name := "hello world"
	fmt.Println(name)
	const langName = "JavaScript"
	fmt.Println(langName)
}

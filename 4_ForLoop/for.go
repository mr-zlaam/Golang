package main

// Note that go contain only one loop which is for loop and it is more than enough for every task which required looping.
func main() {
	// While loop using for loop
	/*
		i := 1
		for i <= 3 {
			fmt.Println(i)
			i = i + 1
		}
	*/
	// ** infinite loop
	/*
		for {
			fmt.Println("Infinite loop")
		}
	*/
	// ** classic for loop in golang
	/*
		for i := 0; i <= 3; i++ {
			if i == 2 {
				continue
			}
			fmt.Println(i)
		}
	*/
	// ** range loop in golang
	/*
		for i := range 11 {
			fmt.Println(i)
		}
	*/
}

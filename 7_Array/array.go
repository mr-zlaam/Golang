package main

import "fmt"

func main() {
	// const length = 4
	// var nums [length]int
	// nums[0] = 1
	// fmt.Println(nums[0])
	// var vals [4]bool
	// fmt.Println(vals)
	// const addingTrue = true
	// const addingFalse = false
	// var name [5]string
	// name[0] = "golang"
	// fmt.Println(name)
	// var declaringArray = [4]int{1, 2, 3, 4}
	// fmt.Println(declaringArray)
	// **  2D Array
	nums := [2][2]int{{3, 4}, {4, 5}}
	fmt.Println(nums)
	/*
		  use Array when you know the fixed size of array  it will optimize memory
			it gives us constant time access
	*/
}

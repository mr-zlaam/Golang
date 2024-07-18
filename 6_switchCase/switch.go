package main

import (
	"fmt"
	"time"
)

func main() {
	// ** Simple Switch Case
	i := 3
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Other")
	}

	// ** Multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Its Saturday")
	case time.Sunday:
		fmt.Println("Its Sunday")
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	}
	//  ** Type Switch
	WhoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other", t)

		}
	}
	WhoamI(true)

	// end
}

package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%v is even \n", i)
	// 	} else {
	// 		fmt.Printf("%v is odd \n", i)
	// 	}
	// }

	n := 5
	for i := range [n]int{} {
		fmt.Println(i)
	}
}

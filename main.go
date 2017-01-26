package main

import "fmt"

func main() {
	for i := 30; i < 250; i++ {
		//	fmt.Println("hello Declan!")
		//	fmt.Println("hello Declan!")
		//fmt.Println(42)
		//fmt.Printf("%d - %b \n", 42, 42)
		fmt.Printf("%d \t %b \t %X \t %q \n", i, i, i, i)
	}
}

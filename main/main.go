package main

import (
	"fmt"
)

func main() {
	println("Hello World From GO!")
	fmt.Printf("%x %b %d %x \n", 72, 72, 72, 72)
	println("\nPrinting out the first 200 characters of a UTF encoder")

	for i := 0; i <= 2000; i++ {
		fmt.Printf("%x %b %d %q \n", i, i, i, i)
	}
}

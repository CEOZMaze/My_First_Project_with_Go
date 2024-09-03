package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	fmt.Printf("%x\n", num)
	fmt.Printf("%X\n", num)
}

// main.go

package main

import "fmt"

func main() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)
}

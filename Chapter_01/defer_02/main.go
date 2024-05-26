// main.go

package main

import "fmt"

func testDefer(val int) {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	if val == 5 {
		return
	}

	defer fmt.Println(4)
}

func main() {
	testDefer(3)
	fmt.Println()
	testDefer(5)
}

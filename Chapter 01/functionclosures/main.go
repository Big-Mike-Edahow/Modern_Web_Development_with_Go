// main.go

package main

import "fmt"

func calc() func() int {
	a := 0
	return func() int {
		fmt.Println("Hello from deep inside the closure!")
		a += 1
		return a
	}
}

func main() {
	res := calc()
	fmt.Println(res())
}

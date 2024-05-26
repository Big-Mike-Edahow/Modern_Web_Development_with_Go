// main.go

package main

import (
	"fmt"
	"hello/countries"
)

func main() {
	fmt.Println("Hello World!!!")

	fmt.Println("Hello", countries.GetCountry("US"))
}

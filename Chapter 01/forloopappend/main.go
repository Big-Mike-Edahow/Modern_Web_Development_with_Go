// main.go

package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3}
	var mySlice []*int

	for _, v := range myArray {
		mySlice = append(mySlice, &v)
	}

	for _, v := range mySlice {
		fmt.Println(*v)
	}
}

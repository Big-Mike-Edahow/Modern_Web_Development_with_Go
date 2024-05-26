// main.go

package main

import "fmt"

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i * 2
	}

	fmt.Println("Index\tElement")
	for index, value := range arr {
		fmt.Printf("%v\t%v\n", index, value)
	}

	sum := 0
	for i, v := range arr {
		if i%2 == 0 {
			sum += v
		}
	}
	fmt.Println("\nTotal of all even indices:", sum)
}

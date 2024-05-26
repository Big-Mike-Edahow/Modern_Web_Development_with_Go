// main.go

package main

import (
	"encoding/json"
	"fmt"
)

type Spouse struct {
	Name string
	Age  int
}

type Person struct {
	Name     string
	Age      int
	Married  bool
	Spouse   Spouse
	Children []string
}

func main() {
	jsonString := `{"name": "Albert Einstein", "age": 76, "married": true, "spouse": {"name": "Mileva", "age": 72}, "children": ["Lieserl", "Hans Albert", "Eduard"]}`
	var p Person
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Failed to decode JSON")
	}
	fmt.Println(p)

	jsonByte, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Failed to encode JSON")
	}
	fmt.Println(string(jsonByte))
}

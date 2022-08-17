package main 

import "fmt"

func main() {
	fmt.Println("Hello, Hashmaps and I run at O(1)")

	m := make(map[string]interface{})

	m["name"] = "Brian Mwangi"
	m["age"] = 24
	m["education"] = "Masters"
	m["work"] = "Azenia"

	fmt.Println("This is my map : ", m)
}

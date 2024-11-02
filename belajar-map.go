package main

import "fmt"

func main() {
	fmt.Println("ganteng")
	person := map[string]string{
		"name" : "ilman",
		"lastname" : "ilman",
	}
	fmt.Println(person)
	person["gender"] = "laki - laki"
	fmt.Println(person)
}
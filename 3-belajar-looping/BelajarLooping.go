package main

import "fmt"

func main() {
	listBilangan := [...]int{
		1, 2, 3, 4, 5,
	}
	for key, value := range listBilangan {
		fmt.Println(key, value)
	}
}
package main

import "fmt"

func getAddition(inputOne int, inputTwo int) int {
	return inputOne + inputTwo
}
func main() {
	resultBilangan := getAddition(2, 3)
	fmt.Println(resultBilangan)
}
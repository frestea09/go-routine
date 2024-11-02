package main

import "fmt"

func main() {
	variableBoolean := true
	variabelInt := 12
	variableString := "hello world"
	variabelArray := []int{1, 2, 3, 4, 5}
	variableSlice := [...]int{0}
	variableSlice[0] = 1
	variabelMap := map[string]string{
		"firstName" : "ilman teguh prasetya",
	}
	fmt.Println(variabelMap)
	fmt.Print(variableSlice)
	fmt.Println(variableBoolean)
	fmt.Println(variabelInt)
	fmt.Println(variableString)
	fmt.Println(variabelArray)
}
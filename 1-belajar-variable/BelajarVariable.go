package main

import "fmt"

func main() {
	fmt.Println("hello world")
	variabelString := "hello world"
	variableInt := 12;
	variableBoolean := true
	variableArray := [...]int{1,2,3,4,5}
	variableMap := map[string]string{
		"firstname" : "ilman teguh prasetya",
	}
	fmt.Println(variabelString)
	fmt.Println(variableInt)
	fmt.Println(variableBoolean)
	fmt.Println(variableArray)
	fmt.Println(variableMap)
}
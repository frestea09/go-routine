package main

import (
	"belajar-go-dasar/5-belajar-struct/moduleofstruct"
	"fmt"
)

func main() {
	tempStruct := moduleofstruct.MyStruct{}
	myString := tempStruct.GetOfStruct()
	fmt.Print(myString)
}
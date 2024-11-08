package main

import (
	"belajar-go-dasar/2-belajar-condition/moduleofcondition"
	"fmt"
)

func main() {
	myImport := moduleofcondition.GetValidation(4)
	fmt.Println(myImport)
}
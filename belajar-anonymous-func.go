package main

import "fmt"

// main is the entry point of the Go program.
//
// It does not take any parameters.
// It does not return any values.

type strucMahasiswa struct{
	name, alamat string
	age int
}
func main() {

	objMashsiswa := strucMahasiswa{
		"ilman",
		"teguh",
		20,
	}
	sayHello := func() int {
		return 12
	}
	hasil := sayHello()
	fmt.Print(hasil)
	fmt.Print(objMashsiswa.age)
}
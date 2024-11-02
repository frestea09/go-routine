package main

import "fmt"

type MahasisStruct struct {
	Nama string
	Age  int
}

func (mahasisprop MahasisStruct) sayHello() {
	fmt.Println("Hello, my name is", mahasisprop.Nama)
}
func main() {
	mahasiswa := MahasisStruct{Nama: "ilman", Age : 23}
	mahasiswa.sayHello()
}
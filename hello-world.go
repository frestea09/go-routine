package main

import "fmt"
func main(){
	var name string;
	type nomorKtp string;
	var ktp nomorKtp = "123456"
	fmt.Println("hello world")
	fmt.Println("Number 1", 1)
	fmt.Println("Dua", 2)
	fmt.Print("tipe boolean", true)
	fmt.Println("ganteng banget")
	fmt.Println(len("ganteng"))
	fmt.Println("ilman teguh"[2])
	name = "ilman teguh prasetya"
	fmt.Println(name)
	variableSatu:=""
	variableSatu = "dua belas"
	fmt.Println("variable satu = ", variableSatu)
	var(
		firstname = "ilman"
		middlename = "teguh"
	)
	fmt.Println("first name", firstname)
	fmt.Println("middle name", middlename)

	const phi = 3.14
	fmt.Println("phi = ", 3.14)

	var variabel32 int32 = 1234
	var variable64 int64 = int64(variabel32)
	fmt.Println(variable64)

	var variableName string = "ilman teguh"
	var ambilVariableName uint8 = variableName[0]
	var estring = string(ambilVariableName)
	fmt.Print(estring)
	fmt.Println("ktp =", ktp)
}
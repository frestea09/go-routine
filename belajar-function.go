package main

import "fmt"

func filter(inputKata string) string {
	if inputKata == "kuda" {
		return ""
	} else {
		return inputKata
	}
}
func sayHello(inputKata string, filter func(string) string) string {
	result := ("hello" + filter(inputKata))
	return result
}
type Mahasiswa struct{
	nim int
	name string
}
func main() {
	blacklist := func(inputName string)string{
		if(inputName=="kuda"){
			return ""
		}else{
			return inputName
		}
		
	}
	resultAnonimus := blacklist("bukan")
	hasil := sayHello("kuda", filter)
	dataMahasiswa := Mahasiswa{
		10,
		"ilman",
	}
	fmt.Print(hasil)
	fmt.Println(resultAnonimus)
	fmt.Println(dataMahasiswa)
}
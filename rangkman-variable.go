package main

import "fmt"

func getAddtion(inputPertama int, inputKedua int) int {
	return inputPertama + inputKedua
}
func getSubstraction(inputPertama int, inputKedua  int) int{
	return inputPertama - inputKedua
}
func getMultiplication(inputPertama int, inputKedua int) int{
	return inputPertama * inputKedua
}
func main() {
	variableString := "variable String"
	variableint := 0
	variableBoolean := true
	variableArray := [...]string{
		"satu",
		"dua",
		"tiga",
	}
	variableMap := map[string]int{
		"satu": 1,
		"dua":  2,
		"tiga": 3,
	}
	fmt.Println(variableString, variableint, variableBoolean, variableArray, variableMap)
	sliceSatu := variableArray[:1]
	const phi = 3.14;
	fmt.Println(sliceSatu)
	fmt.Println(getAddtion(4,5))
	fmt.Println(phi)
}
package main

import "fmt";

func getReturnMulti() (string, string){
	return "hello","ganteng"
}
func getSumAll(inputNumber ...int) int{
	result := 0
	for _,value := range inputNumber{
		result+=value
	}
	return result
}
func main(){
	fmt.Println("hello world")
	listBilangan := [...]int{
		1,2,3,4,5,6,
	}
	for key, value:= range listBilangan{
		fmt.Println("index", key, "=", value)
	}
	for iteration:=0; iteration<=5; iteration++{
		fmt.Println(iteration)
	}
	firtName, lastname := getReturnMulti()
	fmt.Println(firtName, lastname)
	hasilSumAll := getSumAll(1,2,3,4)
	fmt.Println(hasilSumAll)
}
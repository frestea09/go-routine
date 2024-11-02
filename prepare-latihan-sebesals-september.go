package main

import "fmt";

func getAddtion(inputPertama int, inputKedua int) int{
	return inputPertama + inputKedua
}
func getSubstraction(inputPertama int, inputKedua int) int{
	return  inputPertama - inputKedua
}

func getCondition(inputOne int, inputTwo int) bool{
	return inputOne > inputTwo
}
func fungsiPenyediaExtraction(inputOne int, inputTwo int) (int, int){
	return inputOne, inputTwo
}
func main(){
	// fmt.Print("hello world");
	// variableString:="variable string";
	// variableNumber:=5;
	// variableBoolean:=true;
	// fmt.Println(variableBoolean, variableNumber, variableString)
	// resultGetAddtion:=0;
	// resultGetAddtion=getAddtion(1,2);
	// fmt.Println(resultGetAddtion)
	// resultSubtraction := getSubstraction(3,2);
	// fmt.Println(resultSubtraction)
	// resultCondtion := getCondition(3,2);
	// fmt.Println(resultCondtion)
	// arrayPertama := [2]int{1,2};
	// fmt.Println(arrayPertama)
	// arrayKedua := [...]int{1,2,3}
	// fmt.Println(arrayKedua)
	// arrayKedua[0] = 99
	// fmt.Println(arrayKedua)
	// temp := arrayKedua[1:]
	// fmt.Println(temp)
	// arrayUntukDipotong := [...]int{1,2,3,4,5,6,7,8,9,10}
	// sliceAwal := arrayUntukDipotong[2:];
	// fmt.Println(sliceAwal)
	// variableMap :=map[string]string{
	// 	"name" : "ilman",
	// }
	// fmt.Println(variableMap["name"])
	// isLogin := true;
	// if(isLogin){
	// 	fmt.Println("login")
	// }else{
	// 	fmt.Println("logout")
	// }
	// name := "ilman";
	// if(name == "ilman"){
	// 	fmt.Println("benar")
	// }else if name== "tegu"{
	// 	fmt.Println("teguh")
	// }else{
	// 	fmt.Println("salah")
	// }
	// listBilangan := []int{1,2,3,4,5,6}
	// for key, value := range listBilangan{
	// 	fmt.Println("index", key, "=", value)
	// }
	inputOne, inputKedua := fungsiPenyediaExtraction(2, 3);
	fmt.Println(inputOne, inputKedua)
}
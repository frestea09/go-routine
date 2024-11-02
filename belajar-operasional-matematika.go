package main

import (
	"fmt"
)

func main() {
	var variablePertama = 2
	var variableKedua = 3
	var hasilAddtion = variablePertama + variableKedua
	var hasilPengurangan = variablePertama - variableKedua
	fmt.Println("hasil tambah", hasilAddtion)

	var iterasi = 0;
	iterasi+=1
	var unaryaja = 0;
	unaryaja++
	fmt.Println(iterasi)
	fmt.Println(hasilPengurangan)
	fmt.Println(unaryaja)

	var thersholdPertama = 2 
	var therholdKedua = 3
	if(thersholdPertama == therholdKedua){
		fmt.Println("ganteng")
	}else{
		fmt.Println("super ganteng")
	}
	var thresholdOne = 2
	var thresholdTwo = 2
	var hasilPertama bool = thresholdOne == thresholdTwo
	var hasilKedua bool = thersholdPertama != therholdKedua
	var hasilAkhir bool = hasilPertama && hasilKedua
	if(hasilAkhir){
		fmt.Println("benar")
	}else{
		fmt.Println("salah")
	}
	var listBilangan [4]int;
	listBilangan[0] = 2
	listBilangan[1] = 3
	listBilangan[2] = 5
	fmt.Println(listBilangan)
	var listBilanganDua [4]int= [4]int{
		90,80,30,50,
	}
	fmt.Println(listBilanganDua)
	var flexiblelist  = [...]int{
		12,
		24,
		34,
	}
	fmt.Println(flexiblelist)
	arraySaja := [...]int{
		12,
		13,
		14,
		15,
		16,
		17,
		18,
	}
	slice := arraySaja[1:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(append(slice,19))
	listHari := [...]string{
		"senin", 
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(listHari)
	listHari[0] = "Senin Baru"
	fmt.Println(listHari)
	sliceSatu := listHari[:3]
	fmt.Println(sliceSatu)
	sliceSatu[0] = "Senin Super Baru"
	fmt.Println(listHari)
	sliceTiga := listHari[3:]
	fmt.Println(sliceTiga)
	listBaru := append(listHari[:], "langka")
	fmt.Println(listBaru)
	newSlice := make([]string, 3, 5)
	newSlice[0] = "Satu"
	newSlice[1] = "Dua"
	newSlice[2] = "Tiga"
	fmt.Println(newSlice)
	tambahappendtiga := append(newSlice, "tiga")
	fmt.Println(tambahappendtiga)
	fmt.Println(newSlice)
}
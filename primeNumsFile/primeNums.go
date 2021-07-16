package primeNumsFile

import "fmt"

func PrimeNumsFunc(x int) {
	bool:=false
	for i := 2; i < x; i++ {
		if x%i==0{
			bool=true
		}
	}
	if bool{
		fmt.Println("The number is NOT PRIME")
	}else{
		fmt.Println("The number is PRIME")
	}
}
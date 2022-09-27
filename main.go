package main

import (
	"fmt"

	naming "github.com/wghglory/baby-naming/naming"
	naming2 "github.com/wghglory/baby-naming/v2/naming"
	naming3 "github.com/wghglory/baby-naming/v3/naming"
)

func main() {
	fmt.Println("v1 --- A random baby name:", naming.CreateBabyName())

	fmt.Println("v2 --- A random female baby name:", naming2.CreateBabyName(false))
	fmt.Println("v2 --- A random male baby name:", naming2.CreateBabyName(true))

	// ----------------------------below v3------------------------------------
	// CreateBabyName second parameter is uint, so no need to consider negative length
	male, maleErr := naming3.CreateBabyName(true, 4)
	if maleErr == nil {
		fmt.Printf("v3 --- male baby found (min length %v): %v\n", 4, male)
	} else {
		fmt.Println("v3 --- ", maleErr)
	}

	female, femaleErr := naming3.CreateBabyName(false, 3)
	if femaleErr == nil {
		fmt.Printf("v3 --- female baby found (min length %v): %v\n", 3, female)
	} else {
		fmt.Println("v3 --- ", femaleErr)
	}

	baby, err := naming3.CreateBabyName(false, 30)
	if err == nil {
		fmt.Printf("v3 --- female baby found (min length %v): %v\n", 30, baby)
	} else {
		fmt.Println("v3 --- ", err)
	}
}

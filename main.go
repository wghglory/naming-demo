package main

import (
	"fmt"

	naming "github.com/wghglory/baby-naming/naming"
	naming2 "github.com/wghglory/baby-naming/v2/naming"
)

func main() {
	fmt.Println("v1 --- A random baby name:", naming.CreateBabyName())

	fmt.Println("v2 --- A random female baby name:", naming2.CreateBabyName(false))
	fmt.Println("v2 --- A random male baby name:", naming2.CreateBabyName(true))
}

package main

import (
	"fmt"

	naming "github.com/wghglory/baby-naming/naming"
)

func main() {
	name := naming.CreateBabyName()
	fmt.Println("A random baby name:", name)
}

package main

import (
	"cgo/asm/pkg"
	"fmt"
)

func main() {
	fmt.Println(pkg.Id)
	fmt.Println(pkg.Name)
	fmt.Println(pkg.NameData)
}

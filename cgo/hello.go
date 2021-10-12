package main

import "cgo/pkg"

//import "C"

func main() {
	println("hello cgo")
	println(pkg.Id)
	println(pkg.Name)
}

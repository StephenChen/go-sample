// hello.go
package main

//#include <stdio.h>
import "C"
import "cgo/pkg"

func main() {
	C.puts(C.CString("hello c world\n"))
	println("hello cgo")
	println(pkg.Id)
	println(pkg.Name)
}

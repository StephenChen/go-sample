//go:build go1.10

// hello.go
package main

// 思考: main 和 SayHello 是否在同一个 Goroutine 里执行?

/*
// 1 run: go run hello.go, no need hello.c
#include <stdio.h>
static void SayHello1(const char* s)
{
	puts(s);
}

// 2 run: go run ., need hello.c
void SayHello2(const char* s);

// 3 run: go run ., need: hello.h, hello.c
// 4 run: go run ., need: hello.h, hello.cpp
// 5 run: go run ., need: hello.h
#include "hello.h"

// 6 run: go run hello.go
void SayHello6(char* s);

// 7 run: go run hello.go
void SayHello7(_GoString_ s);
*/
import "C"
import "fmt"

//export SayHello5
func SayHello5(s *C.char) {
	fmt.Print(C.GoString(s))
}

//export SayHello6
func SayHello6(s *C.char) {
	fmt.Print(C.GoString(s))
}

//export SayHello7
func SayHello7(s string) {
	fmt.Print(s)
}

func main() {
	C.puts(C.CString("0.hello c world\n"))
	C.SayHello1(C.CString("1.hello c world\n"))
	C.SayHello2(C.CString("2.hello c world\n"))
	C.SayHello3(C.CString("3.hello c world\n"))
	C.SayHello4(C.CString("4.hello c world\n"))
	C.SayHello5(C.CString("5.hello c world\n"))
	C.SayHello6(C.CString("6.hello c world\n"))
	C.SayHello7("7.hello c world\n")
}

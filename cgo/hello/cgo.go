package main

// Point:
//   1. *main.C.char 和 *cgo_helper.C.char 是不同类型。
//   2. 通过 #cgo 设置编译和链接阶段相关参数。
//     CFLAGS 部分:
//       -D 定义了宏,
//       -I 定义了头文件包含的检索目录。
//     LDFLAGS 部分:
//       -L 指定了链接时库文件检索目录，
//       -l 指定了链接时需要链接的库。
//     CPPFLAGS, CXXFLAGS, FFLAGS。

// #cgo CFLAGS: -D PNG_DEBUG=1 -I ./include
// #cgo LDFLAGS: -L /usr/local/lib -l png
// #cgo windows CFLAGS: -D X86=1
// #cgo !windows LDFLAGS: -l m
// #include <png.h>
import "C"

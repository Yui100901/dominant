package main

import (
	"fmt"
	"unsafe"
)

//
// @Author yfy2001
// @Date 2024/7/5 16 21
//

type Programmer struct {
	name     string
	language string
}

func Add(a, b int64) int64 {
	return a + b
}

func Sub(a, b int64) int64 {
	return a - b
}

func main() {
	p := &Programmer{"abc", "go"}
	p1 := Programmer{"134", "c"}
	fmt.Printf("%p\n", Add)
	fmt.Printf("%p\n", Sub)

	name := (*string)(unsafe.Pointer(p))
	*name = "def"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"
	//
	fmt.Println(p)
	fmt.Println(p1)
}

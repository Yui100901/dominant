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

func main() {
	p := &Programmer{"abc", "go"}
	fmt.Printf("%p", p)

	name := (*string)(unsafe.Pointer(p))
	*name = "def"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"
	//
	fmt.Println(p)
}

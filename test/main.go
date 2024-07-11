package main

import "fmt"

//
// @Author yfy2001
// @Date 2024/7/5 16 21
//

func main() {
	u := &User{}
	fmt.Printf("%p %v", u, u.Name)

}

type User struct {
	Name string
}

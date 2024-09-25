package main

import (
	"dominant/domain/node"
	"fmt"
)

func main() {

	n := node.NewNode("a", "1", 3)
	fmt.Println(n.Auth.Verify())
}

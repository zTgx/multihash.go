package main

import (
	"fmt"
	"multihash.go"
)

func main() {
	var ret = multihash.Sm3();
	fmt.Println("ret: ", ret);
}
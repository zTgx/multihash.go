package main

import (
	"fmt"
	"multihash.go"
)

func main() {
	var ret = multihash.Sm2();
	fmt.Println("ret: ", ret);
}
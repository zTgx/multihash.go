package core

import (
	"multihash.go"
)

func Core() string {
	return multihash.Sm3();
}
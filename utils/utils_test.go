package utils

import (
	"multihash.go/core"
	"testing"
)

func TestSm3(t *testing.T)  {
	var target = "sm3"
	var ret = core.Core();
	if ret != target {
		t.Error("sm3 失败。")
	} 
}

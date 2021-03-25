package core
    
import (
	"testing"
	"multihash.go"
)

func TestSm3(t *testing.T)  {
	var target = "sm3"
	var ret = multihash.Sm3()
	if ret != target {
		t.Error("sm3 失败。")
	} 
}

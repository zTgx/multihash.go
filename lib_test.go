package multihash
    
import (
	"testing"
)

func TestSm3(t *testing.T)  {
	var target = "sm3"
	var ret = Sm3()
	if ret != target {
		t.Error("sm3 失败。")
	} 
}

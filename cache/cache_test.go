package cache
import (

	"testing"
)


func TestSimpleCache(t *testing.T) {
	sdc := NewSimpleDiskCache("resource")
	sdc.Set("a_b_c_d_e", []byte("4343"))
	data, _ := sdc.Get("a_b_c_d_e")
	s := "21b528a22f93500928a6fbafe24bbb52"
	t.Log(s[1:] == "png")
	t.Log(s[:8], s[8:16], s[16:24], s[24:])
	t.Log(string(data))
}


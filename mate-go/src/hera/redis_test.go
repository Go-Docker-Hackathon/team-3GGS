package hera

import (
	"fmt"
	"testing"
)

func TestDofunc(t *testing.T) {
	NewRedisSvc()
	tmp, _ := Redis.DoCmd("GET", "uid")
	fmt.Printf("%v", tmp)
}

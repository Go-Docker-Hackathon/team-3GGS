package mate

import (
	"fmt"
	"hera"
	"testing"
)

func TestHelloGet(t *testing.T) {
	url := "http://localhost:8083/Hello/Get?fd=1fda"
	var ret hera.ReturnValue = CurlFunc(url)
	fmt.Printf("ret value:%#v\n", ret)
}

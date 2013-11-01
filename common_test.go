package pcs

import (
	"fmt"
	"testing"
)

func TestReqJson(t *testing.T) {
}

func TestParseUrlValues(t *testing.T) {
	lr := &ListReq{}
	lr.Method = "hello"
	lr.Path = "paa"

	fmt.Printf("%#v\n", parseUrlValues(lr))
}

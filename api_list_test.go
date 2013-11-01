package pcs

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	req := NewListReq()
	req.AccessToken = ACCESS_TOKEN
	req.Path = `/apps/kunyi`

	fmt.Printf("%#v\n", req)
	resp, err := List(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", resp)
}

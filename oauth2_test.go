package pcs

import (
	"fmt"
	"os"
	"testing"
)

func TestReqDeviceCode(t *testing.T) {
	fd, err := os.OpenFile("/dev/fd/1", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	dcr, err := ReqDeviceCode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", dcr)

	var state int64
	fmt.Print("enter a number to continue:")
	n, err := fmt.Fscanf(fd, "%d\n", &state)
	if err != nil {
		fmt.Println(n, err)
		os.Exit(1)
	}
	fmt.Println("state=", state)

	atr, err := ReqAccessToken(dcr.DeviceCode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", atr)
}

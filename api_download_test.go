package pcs

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	req := NewDownloadReq()
	req.AccessToken = ACCESS_TOKEN
	req.Path = `/apps/kunyi/mo_mt_detail.xlsx`

	f, _ := os.Create("abc.xlsx")
	err := Download(req, f)
	if err != nil {
		panic(err)
	}
}

package pcs

import (
	"io"
)

func NewDownloadReq() *BasicFileReq {
	return &BasicFileReq{Method: "download"}
}

func Download(req *BasicFileReq, w io.Writer) (err error) {
	data := parseUrlValues(req)
	err = getData(D_PCS_URL, data, w)
	return
}

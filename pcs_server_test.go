package pcs

import (
	"net/http"
	"testing"
)

func musicHandler(w http.ResponseWriter, r *http.Request) {
	req := NewDownloadReq()
	req.AccessToken = ACCESS_TOKEN
	req.Path = `/apps/kunyi/a.mp3`

	err := Download(req, w)
	if err != nil {
		panic(err)
	}
}

func TestServer(t *testing.T) {
	http.HandleFunc("/music/a.mp3", musicHandler)
	http.ListenAndServe(":8808", nil)
}

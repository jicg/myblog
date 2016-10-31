package comm

import (
	"fmt"
	"net/http"
)

type ImageHandle struct {
}

func (h *ImageHandle) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Println(path)
	rw.Write([]byte("jicg"))
}

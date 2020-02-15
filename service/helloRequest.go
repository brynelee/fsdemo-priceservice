package service

import (
	"fmt"
	"net/http"
)

type PSService struct {
	num int
}

func HelloRequest(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello, this is from Price Service, request url is: ", request.URL.Path[1:])
}

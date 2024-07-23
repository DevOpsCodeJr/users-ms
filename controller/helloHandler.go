package controller

import (
	"fmt"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

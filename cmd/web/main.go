package main

import (
	"fmt"
	"github/satishmd/mock/handler"
	"net/http"
)

const portNumber = ":9001"

func main() {
	http.HandleFunc("/dl", handler.DlWidget)
	fmt.Println(fmt.Sprintf("Application running on %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}

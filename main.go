package main

import (
	"net/http"

	"github.com/praveenksam/goWebServer/api"
)

func main(){
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
